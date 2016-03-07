/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. 
You can only see the k numbers in the window. Each time the sliding window moves right by one position.
*/
package main

import (
    "errors"
    "fmt"
)

func main()  {
    array := []int{1,-1,0,-2,5,-1,7,5}
    r,_ := maxSlidingWindow(array,2) 
    fmt.Printf("%v",r)       
}

func maxSlidingWindow( array []int, k int) ([]int, error) {
    
    if array == nil || len(array)<=0 {
        panic(errors.New("invalid array"))
    }
    count:= len(array)
    if k<=0 || k>count {
        panic(errors.New("invalid k"))
    }
    max := make([]int,count-k+1)
    for i := 0; i < count-k+1; i++ {
        for j := i; j <i+k ; j++ {            
            if array[j] > max[i] {
                max[i]=array[j]
            }
        }
    }    
    return max,nil
}

