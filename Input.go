package main;

import "os"; //package to open outside files and release the file handler

func readInput() (string, error) {
    buffer := make([]byte, 100)

    cnt, err := os.Stdin.Read(buffer)
    if err != nil {
        return "", err
    }

    if cnt == 1 && buffer[0] == 0x1b {
        return "ESC", nil
    }else if cnt >= 3 {
        if buffer[0] == 0x1b && buffer[1] == '[' {
            switch buffer[2] {
            case 'A':
                return "UP", nil
            case 'B':
                return "DOWN", nil
            case 'C':
                return "RIGHT", nil
            case 'D':
                return "LEFT", nil
            }
        }
    }

    return "", nil
}