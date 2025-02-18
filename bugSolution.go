func calculate(a, b int) (int, error) {
  if b == 0 {
    return 0, errors.New("division by zero")
  }
  return a / b, nil
}

func main() {
  result, err := calculate(10, 0)
  if err != nil {
    fmt.Printf("Error: %%v\n", err) //Improved error handling
    return //Exit if error occurs
  } 
  fmt.Println("Result:", result)
} 