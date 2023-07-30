func replaceSpace(s string) string {
    var res string;
    for _,c := range s{
        if( ' '==c ){
            res+="%20";
        }else{
            //下面这样的有问题！下面是char而不是string
            // res+=c;
            res+=string(c);
        }
    }
    return res;
}