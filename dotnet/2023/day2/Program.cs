
using System.Text;
using System.Text.RegularExpressions;

string file = "../../../input.txt";

if (args.Length > 0)
{
    file = args[0];
}

const int red = 12, green = 13, blue = 14;
var games = new List<int>();
var powSum = 0;

using (var reader = new StreamReader(file))
{
    string line;
    int ln = 0;
    while ((line = reader.ReadLine()) != null)
    {
        ln++;
        if (isColorValid(line, "red", red) && isColorValid(line, "green", green) &&
            isColorValid(line, "blue", blue))
        {
            games.Add(ln);
        }

        powSum += getPower(line);
    }
}

Console.WriteLine($"Final sum is {games.Sum()}");
Console.WriteLine($"Powersum is {powSum}");

int getPower(string line)
{
    int hr = 1, hg = 1, hb = 1;
    var rgb = new Dictionary<string, int>();
    rgb.Add("red", 0);
    rgb.Add("green", 0);
    rgb.Add("blue", 0);
    var pattern = @"(?<colorCount>\d+) (?<color>(red|green|blue))";
    var rx = new Regex(pattern);
    MatchCollection matches = rx.Matches(line);
    foreach (Match match in matches)
    {
        var groups = match.Groups;
        rgb.TryGetValue(groups["color"].Value, out var max);
        var curVal = Convert.ToInt32(groups["colorCount"].Value);
        if (curVal > max)
        {
            rgb[groups["color"].Value] = curVal;
        }
    }

    var foo = rgb.Values;
    return foo.Aggregate(1, (x, y) => x * y);
}

bool isColorValid(string line, string color, int max)
{
    var pattern = new StringBuilder()
        .Append(@"(?<colorCount>\d+) ")
        .Append(color)
        .ToString();
    var rx = new Regex(pattern);
    MatchCollection matches = rx.Matches(line);
    foreach (Match match in matches)
    {
        GroupCollection groups = match.Groups;
        if (Convert.ToInt32(groups["colorCount"].Value) > max)
        {
            return false;
        }
    }
    return true;
}