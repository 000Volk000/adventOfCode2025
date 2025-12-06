# frozen_string_literal: true

fich = []
File.foreach('day6.txt') { |line| fich.append(line) }

op_matrix = []

fich.each do |line_complete|
  line = line_complete.split
  op_matrix.append(line)
end

sum_tot = 0
index = op_matrix[0].length - 1
op_matrix[0].each do
  op = ''
  result = 0
  (op_matrix.length - 1).downto(0).each do |row|
    case op
    when ''
      op = op_matrix[row][index]
      result = 1 if op == '*'
    when '+'
      result += op_matrix[row][index].to_i
    when '*'
      result *= op_matrix[row][index].to_i
    end
  end
  sum_tot += result
  index -= 1
end

puts "The total sum of the problems is #{sum_tot}"
