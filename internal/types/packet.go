package types

import "math/rand"

// Packet represents a single data packet moving on the grid
type Packet struct {
	X, Y       int
	PacketType rune // The letter this packet represents
	Dest       int  // Legacy field
	DirX, DirY int  // Movement direction
}

// NewPacket creates a new packet with the specified type
func NewPacket(x, y int, packetType rune) *Packet {
	return &Packet{
		X:          x,
		Y:          y,
		PacketType: packetType,
		Dest:       1, // Legacy
		DirX:       1, // Move right by default
		DirY:       0,
	}
}

// NewRandomPacket creates a packet with a random type from the target word
func NewRandomPacket(x, y int, targetWord string) *Packet {
	if len(targetWord) == 0 {
		return NewPacket(x, y, 'X')
	}

	targetRunes := []rune(targetWord)
	randomType := targetRunes[rand.Intn(len(targetRunes))]

	return NewPacket(x, y, randomType)
}

// Move updates the packet's position based on its direction
func (p *Packet) Move() {
	p.X += p.DirX
	p.Y += p.DirY
}

// SetDirection changes the packet's movement direction
func (p *Packet) SetDirection(dirX, dirY int) {
	p.DirX = dirX
	p.DirY = dirY
}

// GetChar returns the character representation of the packet
func (p *Packet) GetChar() rune {
	return p.PacketType
}

// GetColor returns the appropriate color for this packet type
func (p *Packet) GetColor() string {
	switch p.PacketType {
	case 'G', 'H', 'W', 'C', 'R':
		return ColorCyan
	case 'O', 'I', 'U', 'X', 'A':
		return ColorGreen
	case 'D', 'E', 'T', 'P', 'S':
		return ColorYellow
	case 'M', 'L', 'N', 'F':
		return ColorMagenta
	default:
		return ColorWhite
	}
}

// IsLetterDestination checks if a character is a letter destination
func IsLetterDestination(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
