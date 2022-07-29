import java.util.*;

public class HelloWorld{

   public static void main(String []args){
        ArrayList<String> stringInput = readInput();
        // ArrayList<Integer> input = new ArrayList();

        // for (String str : stringInput) {
        //     input(getIntListFromString(str));
        // }

                // System.out.print("size: "+stringInput.size());


        for (int i = 2; i < stringInput.size(); i = i + 2) {
                            // System.out.print("loop: "+i);

            boolean cont = true;
                // System.out.print("i: "+i);
            ArrayList<Integer> input = getIntListFromString(stringInput.get(i).toString());
            // System.out.print(input.toString());
            int mod = input.get(0) % 2;
            for (int k = 2; k < input.size(); k = k + 2) {
                if (mod != (input.get(k) % 2)) {
                    System.out.println("NO");
                    // return;
                    cont = false;
                    break;
                }
            }
            if (cont) {
                mod = input.get(1) % 2;
                for (int k = 3; k < input.size(); k = k + 2) {
                    if (mod != (input.get(k) % 2)) {
                        System.out.println("NO");
                        // return;
                        cont = false;
                        break;
                    }

                }
                if (cont) {
                    System.out.println("YES");
                }

            }

        }


    }



    public static ArrayList<String> readInput() {
       ArrayList<String> inputList = new ArrayList();
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
          inputList.add(sc.nextLine());
        }
        return inputList;
    }

     public static ArrayList<Integer> getIntListFromString(String str){
      String[] strArray = str.split(" ");
      ArrayList<Integer> intList = new ArrayList();
      for(int i = 0; i < strArray.length; i++){
        intList.add(Integer.parseInt(strArray[i]));
      }
      return intList;
    }
}
