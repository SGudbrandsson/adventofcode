// See https://aka.ms/new-console-template for more information

using System.Collections.Specialized;
using System.Text;

var usage = "USAGE: ./dayN INPUTFILE";
string file = "../../../input.txt";

if (args.Length > 0)
{
    file = args[0];
}
var values = new List<int>();
var values2 = new List<int>();
var numbers = initNumbers();
using (var reader = new StreamReader(file))
{
    string line;
    while ((line = reader.ReadLine()) != null)
    {
        values.Add(ParseLine(line));
        values2.Add(ParseLine(line, true));
    }
}

var total = Enumerable.Sum(values);
var total2 = Enumerable.Sum(values2);
Console.WriteLine($"Total is: {total}");
Console.WriteLine($"Total2 is: {total2}");
return 0;

Dictionary<string, int> initNumbers()
{
    var numbers = new Dictionary<string, int>();
    numbers.Add("one", 1);
    numbers.Add("two", 2);
    numbers.Add("three", 3);
    numbers.Add("four", 4);
    numbers.Add("five", 5);
    numbers.Add("six", 6);
    numbers.Add("seven", 7);
    numbers.Add("eight", 8);
    numbers.Add("nine", 9);
    return numbers;
}

int ParseLine(string input, bool convertStrings = false)
{
    var numList = new SortedList<int, char>();
    int intToCharMagic = 48; // Sum this with an int to get the char value
    if (convertStrings)
    {
        foreach (KeyValuePair<string, int> num in numbers)
        {
            var pos = -1;
            while ((pos = input.IndexOf(num.Key, pos + 1)) != -1)
            {
                numList.TryAdd(pos, (char)(num.Value + intToCharMagic));
            }
        }
    }
    for (int i = 0; i < input.Length; i++)
    {
        if (Char.IsDigit(input[i]))
        {
            numList.TryAdd(i, input[i]);
        }
    }

    if (numList.Count > 0)
    {
        var strNumber = new string([numList.First().Value, numList.Last().Value]);
        
        var baz = Int32.Parse(strNumber);
        return baz;
    }
    return 0;
}

