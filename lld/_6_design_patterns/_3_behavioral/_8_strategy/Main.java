package _6_design_patterns._3_behavioral._8_strategy;

import java.util.Arrays;

// Step 1: Define the Strategy Interface

interface SortingStrategy {
    void sort(int[] numbers);
}

// Step 2: Implement Concrete Strategies

class BubbleSort implements SortingStrategy {
    @Override
    public void sort(int[] numbers) {
        System.out.println("Using Bubble Sort");
        int n = numbers.length;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < n - i - 1; j++) {
                if (numbers[j] > numbers[j + 1]) {
                    int temp = numbers[j];
                    numbers[j] = numbers[j + 1];
                    numbers[j + 1] = temp;
                }
            }
        }
    }
}

class QuickSort implements SortingStrategy {
    @Override
    public void sort(int[] numbers) {
        System.out.println("Using Quick Sort");
        quickSort(numbers, 0, numbers.length - 1);
    }

    private void quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int pi = partition(arr, low, high);
            quickSort(arr, low, pi - 1);
            quickSort(arr, pi + 1, high);
        }
    }

    private int partition(int[] arr, int low, int high) {
        int pivot = arr[high];
        int i = (low - 1);
        for (int j = low; j < high; j++) {
            if (arr[j] <= pivot) {
                i++;
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
        int temp = arr[i + 1];
        arr[i + 1] = arr[high];
        arr[high] = temp;
        return i + 1;
    }
}

// Step 3: Create the Context

class SortContext {
    private SortingStrategy strategy;

    public void setStrategy(SortingStrategy strategy) {
        this.strategy = strategy;
    }

    public void executeSort(int[] numbers) {
        if (strategy == null) {
            throw new IllegalStateException("Strategy not set");
        }
        strategy.sort(numbers);
    }
}

// Step 4: Use the Strategy Pattern

public class Main {
    public static void main(String[] args) {
        SortContext context = new SortContext();

        int[] numbers = {5, 2, 8, 1, 3};

        // Use Bubble Sort
        context.setStrategy(new BubbleSort());
        context.executeSort(numbers);
        System.out.println("Sorted: " + Arrays.toString(numbers));

        // Use Quick Sort
        numbers = new int[]{5, 2, 8, 1, 3}; // Reset the array
        context.setStrategy(new QuickSort());
        context.executeSort(numbers);
        System.out.println("Sorted: " + Arrays.toString(numbers));
    }
}

// Output
// Using Bubble Sort
// Sorted: [1, 2, 3, 5, 8]
// Using Quick Sort
// Sorted: [1, 2, 3, 5, 8]

/** 
Key Points of the Example

    Strategies (BubbleSort, QuickSort): Encapsulate the sorting algorithms.
    Context (SortContext): Uses a strategy to perform the sorting task.
    Flexibility: New sorting strategies (e.g., MergeSort) can be added without changing existing code.

*/