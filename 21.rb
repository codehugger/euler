sum = 0

def d(n)
  (1...((n/2.0).ceil+1)).select { |x| n % x == 0 }.reduce(:+)
end

(1...10000).each do |a|
  b = d(a)
  if a == d(b) and a != b
    sum += a
  end
end

puts sum