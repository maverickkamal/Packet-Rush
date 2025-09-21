package types

import "math/rand"

type Packet struct {
	X, Y       int
	PacketType rune 
	Dest       int  
	DirX, DirY int  
}


func NewPacket(x, y int, packetType rune) *Packet {
	return &Packet{
		X:          x,
		Y:          y,
		PacketType: packetType,
		Dest:       1, 
		DirX:       1, 
		DirY:       0,
	}
}

func NewRandomPacket(x, y int, targetWord string) *Packet {
	if len(targetWord) == 0 {
		return NewPacket(x, y, 'X')
	}

	targetRunes := []rune(targetWord)
	randomType := targetRunes[rand.Intn(len(targetRunes))]

	return NewPacket(x, y, randomType)
}

func (p *Packet) Move() {
	p.X += p.DirX
	p.Y += p.DirY
}

func (p *Packet) SetDirection(dirX, dirY int) {
	p.DirX = dirX
	p.DirY = dirY
}

func (p *Packet) GetChar() rune {
	return p.PacketType
}

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

func IsLetterDestination(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
