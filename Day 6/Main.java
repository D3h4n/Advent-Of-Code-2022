import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) {
        String input;
        String fileName = args[1];
        int length = Integer.parseInt(args[2]);

        try {
            input = readFile(fileName);
        } catch (FileNotFoundException e) {
            System.err.println(e);
            return;
        }
        
        System.out.println("Result: " + findMarkerPosition(input, length));
    }

    static String readFile(String fileName) throws FileNotFoundException {
        StringBuilder input = new StringBuilder();
        File f = new File(fileName);
        Scanner scanner = new Scanner(f);
        
        while(scanner.hasNextLine()) {
            input.append(scanner.nextLine());
        }

        scanner.close();
        return input.toString();
    }

    static int findMarkerPosition(String input, int length) {
        HashSet<String> set = new HashSet<String>();

        if (input.length() < length) 
            return -1;

        for (int i = 0; i < input.length() - length; i++) {
            String[] chars = input.substring(i, i + length).split("");
            Collections.addAll(set, chars);
            
            if (set.size() == length)
                return i + length;

            set.clear();
        }

        return -1;
    }
}
