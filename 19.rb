sundays = 0;

(1901..2000).each do |y|
  (1..12).each do |m|
    sundays += 1 if Time.new(y, m, 1).sunday?
  end
end

puts sundays
