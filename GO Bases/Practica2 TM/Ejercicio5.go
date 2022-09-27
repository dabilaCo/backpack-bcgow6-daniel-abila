package main

func main() {

}

func Animal(animal string) (func(int) int, string) {
  switch animal {
  case "dog":
    return dogFood, ""
  case "cat":
    return catFood, ""
  case "hamster":
    return hamsterFood, ""
  case "tarantula":
    return tarantulaFood, ""    
  default:
    return nil, "No se encontró el animal '" + animal + "'."
  }
}

func dogFood(count int) int { return count * 10 * 1000 }

func catFood(count int) int { return count * 5 * 1000 }

func hamsterFood(count int) int { return count * 250 }

func tarantulaFood(count int) int { return count * 150 }