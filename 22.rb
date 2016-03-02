# Using names.txt (right click and 'Save Link/Target As...'), a 46K text
# file containing over five-thousand first names, begin by sorting it
# into alphabetical order. Then working out the alphabetical value for
# each name, multiply this value by its alphabetical position in the
# list to obtain a name score.

# For example, when the list is sorted into alphabetical order, COLIN,
# which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
# So, COLIN would obtain a score of 938 × 53 = 49714.

# What is the total of all the name scores in the file?

@char_map = {}
sum = 0

("A".."Z").each_with_index do |char, i|
  @char_map[char] = i + 1
end

def char_sum(name)
  name.split("").map{ |c| @char_map[c] }.reduce(:+)
end

File.open("assets/22.txt", "r") do |f|
  f.each_line do |line|
    names = line.split(",").sort

    names.each_with_index do |name, i|
      puts "#{i+1} #{name} -> #{char_sum(name.tr('"', '')) * (i+1)}"
      sum += (char_sum(name.tr('"', '')) * (i+1))
    end
  end
end

puts "Total: #{sum}"