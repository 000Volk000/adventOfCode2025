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

  public static Pair getBiggerCharPos(String line) {
    Pair solution = new Pair(0, line.charAt(0));

    for (int i = 1; i < line.length(); i++) {
      if (line.charAt(i) > solution.getValue()) {
        solution.setKey(i);
        solution.setValue(line.charAt(i));
      }
    }

    return solution;
  };

  public static Pair getBiggerCharPosInverted(String line) {
    Pair solution = new Pair(0, line.charAt(0));

    for (int i = line.length() - 1; i > 0; i--) {
      if (line.charAt(i) > solution.getValue()) {
        solution.setKey(i);
        solution.setValue(line.charAt(i));
      }
    }

    return solution;
  };

  public static void main(String[] args) {
    File fich = new File("day3-example.txt");

    try (Scanner reader = new Scanner(fich)) {
      Long sumTot = Long.parseLong("0");

      while (reader.hasNextLine()) {
        String line = reader.nextLine();
        ArrayList<Pair> numbers = new ArrayList<>();
        Integer erased = 0;

        for (int i = 0; i < 12; i++) {
          Pair s = new Pair(getBiggerCharPos(line));
          System.out.println(
              "skey: " + s.getKey() + "lleng: " + line.length() + "nsiz: " + numbers.size() + "eras: " + erased);
          if ((s.getKey() < line.length() - (11 - numbers.size())) && (erased < (line.length() - 12))) {
            String newLine = "";
            for (int j = 0; j < s.getKey() + 1; j++) {
              newLine = newLine + (char) ('0' - 1);
            }
            line = newLine + line.substring(s.getKey() + 1, line.length());
          } else {
            s = new Pair(getBiggerCharPosInverted(line));
            line = line.substring(0, s.getKey()) + (char) ('0' - 1) + line.substring(s.getKey() + 1, line.length());
          }
          numbers.add(s);
          erased++;
          System.out.println(line);
        }

        numbers.sort((a, b) -> {
          return Integer.compare(a.getKey(), b.getKey());
        });

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
