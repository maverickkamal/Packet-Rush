# 🎮 **Packet Rush - Complete Gameplay Guide**

## 🚀 **Quick Start**

```bash
# Build and run the game
go run main.go

# Or build executable first
go build -o packet-rush.exe .
./packet-rush.exe
```

## 🎯 **Game Overview**

You are a **network router administrator**! Your job is to route data packets to their correct destinations by switching junction directions in real-time.

## 📊 **What You'll See**

### **The Game Grid:**
```
################################################################################
#S----+----+----1                                                             #
#     |    |                                                                  #
#     |    +----2                                                             #
#     |                                                                       #
#     +----+----3                                                             #
#          |                                                                  #
#          +----4                                                             #
```

- **`S`** = Spawn point (where packets appear)
- **`+`** = Junctions (controllable with keys 1-4)
- **`1-4`** = Destination ports
- **`-` `|`** = Network cables/paths
- **`#`** = Walls

### **Moving Elements:**
- **`D`** = Data packet
- **`A`** = Audio packet  
- **`V`** = Video packet
- **`E`** = Email packet

### **Junction States:**
- **`>`** = Junction routing RIGHT
- **`v`** = Junction routing DOWN
- **`^`** = Junction routing UP
- **`<`** = Junction routing LEFT

## 🎮 **Controls**

| Key | Action |
|-----|--------|
| **1-4** | Switch junction directions |
| **SPACE** | Pause/Resume game |
| **Q** | Quit game |
| **R** | Restart (when game over) |

## 🎯 **How to Play**

### **Step 1: Watch Packets Spawn**
- Packets appear at `S` every few seconds
- Each packet is a letter: `D`, `A`, `V`, or `E`
- Each packet has a random destination (port 1, 2, 3, or 4)

### **Step 2: Route Packets Correctly**
- Press keys **1-4** to switch junction directions
- Watch the `+` symbols change to arrows: `>`, `v`, `^`, `<`
- Route packets to their correct numbered destinations

### **Step 3: Score Points**
- **+10 points** for each correctly routed packet
- **-1 life** for wrong deliveries or lost packets
- **Game over** when you lose all 3 lives ("KERNEL PANIC!")

## 📈 **Difficulty Progression**

- **Packets spawn faster** as time progresses
- **Multiple packets** can be on screen simultaneously
- **Junction timing** becomes more critical
- **Quick decisions** are essential for high scores

## 🏆 **Pro Tips**

1. **Plan ahead**: Watch where multiple packets are going
2. **Switch quickly**: Junctions affect the next packet to reach them
3. **Count destinations**: Learn the port numbers (1-4)
4. **Use pause**: Press SPACE to plan your moves
5. **Practice patterns**: Learn common routing sequences

## 🎖️ **Scoring System**

- **Correct delivery**: +10 points
- **Wrong delivery**: -1 life
- **Lost packet**: -1 life
- **Starting lives**: 3
- **Goal**: Survive as long as possible!

## 🔧 **Technical Details**

- **Built with**: Go + Bubble Tea terminal framework
- **Real-time**: 60fps smooth gameplay
- **Cross-platform**: Runs on Windows, Mac, Linux
- **Single binary**: No dependencies needed

---

**🎉 Ready to become the ultimate network router administrator? Run the game and start routing packets!**

