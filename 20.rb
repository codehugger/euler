def fact(n)
  if n == 0
    1
  else
    n * fact(n-1)
  end
end

puts fact(100).to_s.each_char.reduce(0) { |sum, c| sum + c.to_i }
