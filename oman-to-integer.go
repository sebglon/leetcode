
func romanToInt(s string) int {
    result :=0
    previous := 0
    
    size:=len(s)
    if size < 1 || size > 15 {
        return result
    }
    chars := []rune(s)
    for i:=size-1;i>=0; i-- {
        newValue := convert(chars[i])
        if newValue==0{
            return 0
        }    
    
        if  newValue>=previous {
            result+=newValue
        } else {
            result-=newValue
        }
        previous=newValue
    }
    return result
}

func convert(char rune) (result int) {
    
            switch char {
            case 'M':
                result+=1000
            case 'D':
                result+=500
            case 'C':
                result+=100
            case 'L':
                result+=50
            case 'X':
                result+=10
            case 'V':
                result+=5
            case 'I':
                result+=1
            default:
                result = 0
        }
    return
}
