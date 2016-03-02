# A perfect number is a number for which the sum of its proper divisors is
# exactly equal to the number. For example, the sum of the proper divisors
# of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect
# number.

# A number n is called deficient if the sum of its proper divisors is less
# than n and it is called abundant if this sum exceeds n.

# As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
# number that can be written as the sum of two abundant numbers is 24. By
# mathematical analysis, it can be shown that all integers greater than 28123
# can be written as the sum of two abundant numbers. However, this upper limit
# cannot be reduced any further by analysis even though it is known that the
# greatest number that cannot be expressed as the sum of two abundant numbers
# is less than this limit.

# Find the sum of all the positive integers which cannot be written as the
# sum of two abundant numbers.

def sum_of_proper_divisors(n)
  (1...((n/2.0).ceil+1)).select { |x| n % x == 0 }.reduce(:+)
end

def perfect_number(n)
  sum_of_proper_divisors(n) == n
end

def deficient(n)
  sum_of_proper_divisors(n) < n
end

def abundant(n)
  sum_of_proper_divisors(n) > n
end

puts "Perfect 28: #{perfect_number(28)}"
puts "Perfect 13: #{perfect_number(13)}"
puts "Abundant 12: #{abundant(12)}"
puts "Abundant 13: #{abundant(13)}"
puts "Deficient 12: #{deficient(12)}"
puts "Deficient 13: #{deficient(13)}"

# (1...28123).each do |i|
#   puts i if deficient(i)
# end