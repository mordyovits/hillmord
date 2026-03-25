package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gliderlabs/ssh"
	"hillmord/game"
)

// lineReader handles line-buffered input with echo for SSH PTY sessions.
// It reads raw bytes from the SSH channel one at a time, handles backspace
// (without allowing deletion past the prompt), echoes characters, and only
// delivers complete lines (terminated by \n) to the caller.
type lineReader struct {
	raw    io.Reader
	writer io.Writer
	buf    []byte // current line being edited
	ready  []byte // completed line waiting to be consumed by Read
}

func (lr *lineReader) Read(p []byte) (int, error) {
	// If we have a completed line ready, deliver it
	if len(lr.ready) > 0 {
		n := copy(p, lr.ready)
		lr.ready = lr.ready[n:]
		return n, nil
	}

	// Read raw input byte-by-byte until we get a full line
	one := make([]byte, 1)
	for {
		_, err := lr.raw.Read(one)
		if err != nil {
			return 0, err
		}
		switch one[0] {
		case '\r', '\n':
			// Enter: echo CR+LF, deliver the line
			lr.writer.Write([]byte("\r\n"))
			line := append(lr.buf, '\n')
			lr.buf = nil
			n := copy(p, line)
			if n < len(line) {
				lr.ready = line[n:]
			}
			return n, nil
		case 127, '\b':
			// Backspace: only delete if there are typed characters
			if len(lr.buf) > 0 {
				lr.buf = lr.buf[:len(lr.buf)-1]
				lr.writer.Write([]byte("\b \b"))
			}
		case 3:
			// Ctrl+C
			return 0, io.EOF
		default:
			// Normal character: echo and buffer
			lr.buf = append(lr.buf, one[0])
			lr.writer.Write(one)
		}
	}
}

// crlfWriter converts \n to \r\n in output, needed for SSH PTY sessions
type crlfWriter struct {
	writer io.Writer
}

func (w *crlfWriter) Write(p []byte) (int, error) {
	// Replace bare \n with \r\n for proper terminal rendering
	var out []byte
	for i := 0; i < len(p); i++ {
		if p[i] == '\n' && (i == 0 || p[i-1] != '\r') {
			out = append(out, '\r', '\n')
		} else {
			out = append(out, p[i])
		}
	}
	_, err := w.writer.Write(out)
	if err != nil {
		return 0, err
	}
	return len(p), nil // return original length so callers aren't confused
}

func main() {
	sshMode := flag.Bool("ssh", false, "Run as SSH server instead of interactive game")
	sshAddr := flag.String("addr", "0.0.0.0:2222", "SSH server address (host:port)")
	flag.Parse()

	if *sshMode {
		// Run SSH server
		if err := RunSSHServer(*sshAddr); err != nil {
			fmt.Fprintf(os.Stderr, "SSH Server error: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Run interactive game
		fmt.Println(game.TitleScreen())

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What shall we call you, brave fool? > ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = "Mudrick the Uncertain"
		}

		g := game.New(name)
		g.Run(reader)
	}
}

// RunSSHServer starts an SSH server on the given address with no authentication required
func RunSSHServer(addr string) error {
	s := &ssh.Server{
		Addr: addr,
		Handler: func(session ssh.Session) {
			handleGameSession(session)
		},
		// Allow any password (no authentication check)
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			return true // Accept any password
		},
		// Also allow public key authentication with any key
		PublicKeyHandler: func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true // Accept any public key
		},
	}

	log.Printf("🏔️ HILLMORD SSH Server starting on %s (no authentication required)", addr)
	log.Printf("Connect with: ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null localhost:2222")
	return s.ListenAndServe()
}

// handleGameSession manages a single game session for an SSH connection
func handleGameSession(session ssh.Session) {
	defer session.Close()

	// Get player name from environment or SSH command
	playerName := session.User()
	if playerName == "" || playerName == "root" {
		playerName = "Mudrick the Uncertain"
	}

	// Wrap output with CR+LF conversion for proper terminal rendering
	out := &crlfWriter{writer: session}

	// Wrap the session with a line reader that handles echo and backspace
	lr := &lineReader{raw: session, writer: session}

	// Show title screen
	io.WriteString(out, game.TitleScreen())

	// Ask for name
	io.WriteString(out, "What shall we call you, brave fool? > ")

	// Create a reader from the line-editing reader
	reader := bufio.NewReader(lr)
	name, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			log.Printf("Error reading name: %v", err)
		}
		return
	}

	name = strings.TrimSpace(name)
	if name == "" {
		name = playerName
		if name == "" {
			name = "Mudrick the Uncertain"
		}
	}

	// Create and run the game with CR+LF output
	g := game.New(name)
	g.RunWithWriter(reader, out)
}
