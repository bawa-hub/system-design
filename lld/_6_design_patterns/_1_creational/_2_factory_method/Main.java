package _6_design_patterns._1_creational._2_factory_method;

interface Food {
    void prepare();
}

class Pasta implements Food {
    public void prepare() {
        System.out.println("Preparing Pasta (Italian Dish)!");
    }
}

class Noodles implements Food {
    public void prepare() {
        System.out.println("Preparing Noodles (Chinese Dish)!");
    }
}

class Curry implements Food {
    public void prepare() {
        System.out.println("Preparing Curry (Indian Dish)!");
    }
}

interface FoodFactory {
    Food createFood();
}

class ItalianFoodFactory implements FoodFactory {
    public Food createFood() {
        return new Pasta(); // Italian factory creates Pasta
    }
}

class ChineseFoodFactory implements FoodFactory {
    public Food createFood() {
        return new Noodles(); // Chinese factory creates Noodles
    }
}

class IndianFoodFactory implements FoodFactory {
    public Food createFood() {
        return new Curry(); // Indian factory creates Curry
    }
}

public class Main {
    public static void main(String[] args) {
        // Create an Italian Food Factory
        FoodFactory italianFactory = new ItalianFoodFactory();
        Food italianFood = italianFactory.createFood();
        italianFood.prepare(); // Output: Preparing Pasta (Italian Dish)!

        // Create a Chinese Food Factory
        FoodFactory chineseFactory = new ChineseFoodFactory();
        Food chineseFood = chineseFactory.createFood();
        chineseFood.prepare(); // Output: Preparing Noodles (Chinese Dish)!

        // Create an Indian Food Factory
        FoodFactory indianFactory = new IndianFoodFactory();
        Food indianFood = indianFactory.createFood();
        indianFood.prepare(); // Output: Preparing Curry (Indian Dish)!
    }
}

/**
 * 
 * What’s Happening Here?

    Each factory specializes in making one type of object. For example:
        ItalianFoodFactory only makes Pasta.
        ChineseFoodFactory only makes Noodles.
        IndianFoodFactory only makes Curry.

    Your application doesn’t directly create objects. Instead, it uses the appropriate factory.

    It’s easy to add new types of food or cuisines:
        Add a new Food class (e.g., Pizza).
        Add a new Factory class (e.g., PizzaFactory).
 */
