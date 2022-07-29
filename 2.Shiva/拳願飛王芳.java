import java.util.Scanner;

public class Main {

  public static final void main(String[] args) {
    Scanner sc = new Scanner(System.in);
    int count = Integer.valueOf(sc.nextLine());
    StringBuilder sb = new StringBuilder();
    if(sc.hasNextLine()) {
      for(int i = 0; i < count; i++) {
        sb.append(calc(sc));
      }
    }
    sc.close();
    System.out.println(sb);
  }

  private static String calc(Scanner sc) {
    double inputCount = Integer.valueOf(sc.nextLine());
    String[] input = sc.nextLine().trim().split(" ");

    int oddCount = (int)Math.ceil(inputCount / 2);
    int evenCount = (int)Math.floor(inputCount / 2);

    int[] arrayOne = new int[oddCount];
    int[] arrayTwo = new int[evenCount];

    Boolean oneIsOdd = null;
    int oneX = 0;
    Boolean twoIsOdd = null;
    int twoX = 0;


    for(int i = 0; i < input.length; i++) {
      // System.out.println(oddCount + "/" + evenCount);

      //Even
      if(i % 2 == 0) {
        arrayOne[oneX] = Integer.valueOf(input[i]);
        oneX++;
      } else {
        arrayTwo[twoX] = Integer.valueOf(input[i]);
        twoX++;
      }
    }

    boolean ansOne = true;
    for(int i = 0; i < arrayOne.length; i++) {
      int x = arrayOne[i];
      if(oneIsOdd == null) {
        oneIsOdd = x % 2 == 1;
        continue;
      }
      if(oneIsOdd) {
        if(x % 2 == 1) {
          continue;
        } else {
          ansOne = false;
          break;
        }
      } else {
        if(x % 2 == 0) {
          continue;
        } else {
          ansOne = false;
          break;
        }
      }
    }

    boolean ansTwo = true;
    for(int i = 0; i < arrayTwo.length; i++) {
      int x = arrayTwo[i];
      if(twoIsOdd == null) {
        twoIsOdd = x % 2 == 1;
        continue;
      }
      if(twoIsOdd) {
        if(x % 2 == 1) {
          continue;
        } else {
          ansTwo = false;
          break;
        }
      } else {
        if(x % 2 == 0) {
          continue;
        } else {
          ansTwo = false;
          break;
        }
      }
    }

  return (ansOne && ansTwo)?"YES\n":"NO\n";
  }

}