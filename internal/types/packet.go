package types

import (
	"fmt"
	"math/rand"
)

// PacketType represents different types of network packets
type PacketType rune

const (
	DataPacket  PacketType = 'D' // Data packet
	AudioPacket PacketType = 'A' // Audio packet
	VideoPacket PacketType = 'V' // Video packet
	EmailPacket PacketType = 'E' // Email packet
)

// Packet represents a single data packet moving on the grid
type Packet struct {
	Pos         Position   // Current position on the grid
	Type        PacketType // Type of packet (D, A, V, E)
	Destination int        // Target port number (1-4)
	Direction   Position   // Current movement direction
	Speed       int        // Movement speed (ticks between moves)
	TickCount   int        // Internal counter for movement timing
	Color       string     // Color for terminal display
}

// NewPacket creates a new packet with random type and destination
func NewPacket(startPos Position, direction Position) *Packet {
	types := []PacketType{DataPacket, AudioPacket, VideoPacket, EmailPacket}
	packetType := types[rand.Intn(len(types))]

	packet := &Packet{
		Pos:         startPos,
		Type:        packetType,
		Destination: rand.Intn(4) + 1, // Random destination 1-4
		Direction:   direction,
		Speed:       1, // Default speed
		TickCount:   0,
		Color:       getColorForPacketType(packetType),
	}

	return packet
}

// Move updates the packet's position based on its direction and speed
func (p *Packet) Move() {
	p.TickCount++
	if p.TickCount >= p.Speed {
		p.Pos = p.Pos.Add(p.Direction)
		p.TickCount = 0
	}
}

// SetDirection changes the packet's movement direction
func (p *Packet) SetDirection(newDir Position) {
	p.Direction = newDir
}

// GetChar returns the character representation of the packet
func (p *Packet) GetChar() rune {
	return rune(p.Type)
}

// String returns a string representation of the packet
func (p *Packet) String() string {
	return fmt.Sprintf("Packet{Type: %c, Dest: %d, Pos: %s, Dir: %s}",
		p.Type, p.Destination, p.Pos, p.Direction)
}

// getColorForPacketType returns the color code for different packet types
func getColorForPacketType(pType PacketType) string {
	switch pType {
	case DataPacket:
		return "cyan"
	case AudioPacket:
		return "green"
	case VideoPacket:
		return "yellow"
	case EmailPacket:
		return "magenta"
	default:
		return "white"
	}
}

// IsAtDestination checks if the packet has reached its destination port
func (p *Packet) IsAtDestination(char rune) bool {
	// Check if the character at current position is a destination port
	// Ports are numbered 1-4
	switch char {
	case '1', '2', '3', '4':
		portNum := int(char - '0')
		return portNum == p.Destination
	default:
		return false
	}
}

