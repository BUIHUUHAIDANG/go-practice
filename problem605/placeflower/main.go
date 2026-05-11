package main

func canPlaceFlowers(flowerbed []int, n int) bool {
    i:= 0
	for i<len(flowerbed) {
		if flowerbed[i]==1{
			i++
			continue
		}else if len(flowerbed)==1 && flowerbed[0]!= 1{
			flowerbed[i] = 1
			n--
		}else if (i == 0  && flowerbed[i+1] != 1 && flowerbed[i]!=1) || (i == len(flowerbed)-1 && flowerbed[i-1]!= 1 &&flowerbed[i]!=1) ||
		(i != 0 && i!= len(flowerbed)-1 &&flowerbed[i]!=1 && flowerbed[i+1] != 1 && flowerbed[i-1]!= 1){
			flowerbed[i] = 1
			n--
		}
		i++
	}
	if n > 0{
		return false 
	}
	return true
}

func main(){
	flowerbed := []int{1}
	n:= 0
	print(canPlaceFlowers(flowerbed,n))
}
