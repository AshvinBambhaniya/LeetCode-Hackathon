func longestCommonPrefix(strs []string) string {
    var prefix string

    temp1 := strings.Split(strs[0], "")
    temp2 := strings.Split(strs[1], "")
    temp3 := strings.Split(strs[2], "")

    for i:=0;i<len(temp1);i++{
        if len(temp2) == i || len(temp3) == i{
            break
        }
        if temp1[i] == temp2[i] && temp1[i] == temp3[i]{
            prefix += temp1[i]
        } 
    }

    return prefix
}
