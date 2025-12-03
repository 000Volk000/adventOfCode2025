import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class day3_2 {
  static class Pair {
    private Integer key;
    private Character value;

    public Pair(Integer key, Character value) {
      this.key = key;
      this.value = value;
    }

    public Pair(Pair pair) {
      this.key = pair.getKey();
      this.value = pair.getValue();
    }

    public Integer getKey() {
      return key;
    }

    public Character getValue() {
      return value;
    }

    public void setKey(Integer key) {
      this.key = key;
    }

    public void setValue(Character value) {
      this.value = value;
    }

    public String toString() {
      String pair = "Key: " + this.key + ", Value: " + this.value;
      return pair;
    }
  }

  public static Pair getBiggerCharPos(String line, Integer start, Integer end) {
    Pair solution = new Pair(start, line.charAt(start));

    for (int i = start; i <= end; i++) {
      if (line.charAt(i) > solution.getValue()) {
        solution.setKey(i);
        solution.setValue(line.charAt(i));
      }
    }

    return solution;
  };

  public static void main(String[] args) {
    File fich = new File("day3.txt");

    try (Scanner reader = new Scanner(fich)) {
      Long sumTot = Long.parseLong("0");

      while (reader.hasNextLine()) {
        String line = reader.nextLine();
        ArrayList<Pair> numbers = new ArrayList<>();
        Integer numSize = 12;
        Integer start = 0;
        Integer end = line.length() - numSize;

        for (int i = 0; i < numSize; i++) {
          Pair s = getBiggerCharPos(line, start, end);
          numbers.add(s);
          start = s.getKey() + 1;
          end++;
        }

        String numberStr = "";
        for (Pair pair : numbers) {
          numberStr = numberStr + pair.getValue();
        }
        sumTot += Long.parseLong(numberStr);
      }

      System.out.println("The total sum of joltage is " + sumTot);
    } catch (FileNotFoundException e) {
      System.out.println("File not found error ocurred");
      e.printStackTrace();
    }
  }
}
