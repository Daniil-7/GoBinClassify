package GoBinClassify
import (
 "math"
 "errors"
)


/*
CoutQustion accepts the number of groups you want to get.
The number should be:
1. The whole
2. More than 1
If the transfer is successful, it returns the number of responses that you must pass to the Answer method.
*/
func CoutQustion(ngroups int) (int, error) {
 if ngroups < 2 {
  err := errors.New("The number must be greater than 2")
  return 0, err
 }
 qustion := math.Log2(float64(ngroups))
 qustion = math.Ceil(qustion)
 return int(qustion), nil
}


/*
CoutGroups accepts the number of questions. Returns an array with a range of the number of groups.
The first element is the minimum number of groups for such a number of questions, the second is the maximum.
The minimum transmitted number is 1.
*/
func CoutGroups(nqustion int) ([]int, error) {
 if nqustion < 1 {
  err := errors.New("The number must be greater than 1")
  return []int{}, err
 }
 nmin := int(math.Pow(2.0, float64(nqustion - 1))) + 1
 nmax := int(math.Pow(2.0, float64(nqustion)))
 return []int{nmin, nmax}, nil
}


/*
Answer(x, y)
The Answer method accepts: number - number of groups and array - answers to questions. 
Array Criteria:
1. The length of the array must be equal to CoutQustion(x) - the number returned by the CoutQustion method with the first argument of the Answer method passed. 
2. The answers must be binary (0 or 1)
*/
func Answer(ngroups int, outarr []int) (int, error) {
 qustion, err := CoutQustion(ngroups)
 if err != nil {
  return 0, err
 }
 if qustion !=  len(outarr) {
  err := errors.New("The number of elements in the array does not match the returned number of the CoutQustion method.")
  return 0, err
 }
 low := 1
 high := ngroups
 for _, el := range outarr {
  mid := int((low + high) / 2)
  switch el {
   case 0: high = mid
   case 1: low = mid
   default:
    err := errors.New("One of the array elements is not binary (not 1, not 0)")
    return 0, err
  }
 }
 return high, nil
}
