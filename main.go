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

	// Show title screen - use io.WriteString which flushes properly
	io.WriteString(session, game.TitleScreen())

	// Ask for name
	io.WriteString(session, "What shall we call you, brave fool? > ")

	// Create a reader from the session
	// The session implements io.Reader/io.Writer from the Channel interface
	reader := bufio.NewReader(session)
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

	// Create and run the game
	g := game.New(name)
	g.RunWithWriter(reader, session)
}
