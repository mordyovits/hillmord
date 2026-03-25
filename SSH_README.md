# 🏔️ HILLMORD SSH Server

Hillmord can now be played over SSH with **no authentication required** and support for **multiple concurrent players**.

## Quick Start

### Start the SSH Server

```bash
./hillmord -ssh -addr 0.0.0.0:2222
```

This starts an SSH server on port 2222 (default) that accepts connections with no authentication.

### Connect as a Player

From any machine on the network:

```bash
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null localhost:2222
```

Or with a username:

```bash
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null testuser@localhost -p 2222
```

Each player gets their own independent game instance.

## Command-Line Options

```bash
./hillmord -ssh -addr <host:port>
```

- `-ssh`: Enable SSH server mode (default: false, runs interactive game)
- `-addr`: SSH server bind address (default: `0.0.0.0:2222`)

### Examples

- **Listen on all interfaces, port 2222:**
  ```bash
  ./hillmord -ssh -addr 0.0.0.0:2222
  ```

- **Listen only on localhost, port 3000:**
  ```bash
  ./hillmord -ssh -addr localhost:3000
  ```

- **Listen on a specific IP:**
  ```bash
  ./hillmord -ssh -addr 192.168.1.100:2222
  ```

## Authentication

The SSH server has **no authentication requirements**:
- ✅ Any username is accepted
- ✅ Any password is accepted  
- ✅ Any SSH public key is accepted

This makes it simple to set up a shared game server where multiple players can jump in without passwords.

## Multiple Concurrent Players

The SSH server uses goroutines to handle each connection independently. You can have dozens of players connected simultaneously, each with their own game state and character progress.

Each player:
- Creates a new game instance on connection
- Has independent character state
- Cannot see or interact with other players' games

## Game Features Over SSH

All HILLMORD features work perfectly over SSH:
- ⚔️ Combat
- 💰 Markets & trading
- 🗺️ Exploration
- 📊 Character stats
- 🎒 Inventory management
- ✨ XP and leveling

## Interactive Mode (Default)

Without the `-ssh` flag, hillmord runs as a normal interactive terminal game:

```bash
./hillmord
```

## SSH Tips

### For macOS/Linux Users

If you get host key verification warnings, use these SSH flags to bypass them:

```bash
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
    -p 2222 anyuser@your-server.com
```

### Batch Connections for Testing

Test multiple concurrent connections:

```bash
for i in {1..5}; do
  timeout 30 ssh -o StrictHostKeyChecking=no \
    -o UserKnownHostsFile=/dev/null -p 2222 player$i@localhost &
done
wait
```

### Keep Alive

If your connection drops frequently over unstable networks:

```bash
ssh -o ServerAliveInterval=30 -o ServerAliveCountMax=3 \
    -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
    -p 2222 anyuser@localhost
```

## Architecture

### SSH Server Implementation

- **Library**: [gliderlabs/ssh](https://github.com/gliderlabs/ssh) - A pure Go SSH server
- **Connection Handling**: Each SSH session spawns a new `Game` instance in its own goroutine
- **I/O**: The SSH session implements `io.Reader` and `io.Writer`, allowing the game to use its normal I/O patterns

### Game Modifications

The game engine was minimally modified to support custom output writers:

1. Added `output io.Writer` field to the `Game` struct
2. Created `RunWithWriter(reader, writer)` method for SSH connections
3. Added helper methods (`printf`, `println`, `print`) that write to the custom output writer
4. All game output now uses these helpers instead of direct `fmt` calls

This architecture means:
- ✅ The game logic remains unchanged
- ✅ Both interactive and SSH modes use the same game engine
- ✅ New features automatically work in both modes

## Building from Source

```bash
go mod tidy
go build -o hillmord
```

## Performance Notes

- Each player connection uses minimal resources (one goroutine per connection)
- Game state is independent per player (no shared mutable state)
- The SSH server can handle hundreds of concurrent connections
- Host key is automatically generated on startup (regenerated each run unless configured otherwise)

For production use, generate a persistent host key:

```bash
# Generate once
ssh-keygen -t rsa -f ~/.ssh/hillmord_key -N ""

# Use it
./hillmord -ssh -addr 0.0.0.0:2222 &
```

(Note: This would require modifications to the SSH server setup to use a custom host key)

## Troubleshooting

**"Connection refused"**
- The server might not be running, or you're connecting to the wrong port
- Check the server's startup message for the correct address

**"Permission denied (publickey,keyboard-interactive)"**
- This shouldn't happen with no auth enabled, but try updating your SSH client

**SSH key warnings with `-o StrictHostKeyChecking=no`**
- This is expected for a new SSH server without a persistent host key
- The warnings are harmless for a game server

## Related Documentation

See [README.md](README.md) for:
- How to play HILLMORD
- Game mechanics and combat
- World map and locations
- Character progression

---

**Happy adventuring!** 🗡️ May your SSH connections be stable and your character's HP remain positive.
