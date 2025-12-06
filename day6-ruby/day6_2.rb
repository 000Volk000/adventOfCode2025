# frozen_string_literal: true

fich = []
File.foreach('day6.txt') { |line| fich.append(line) }

op_matrix = []

fich.each do |line_complete|
  op_matrix.append(line_complete)
end

sum_tot = 0
index = op_matrix[0].length - 1
num = ''
num_vec = []
op_matrix[0].length.times do
  op = ''
  result = 0

  op_matrix.length.times do |row|
    case op_matrix[row][index]
    when '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
      num += op_matrix[row][index]
    when '+', '*'
      op = op_matrix[row][index]
    end
  end

  num_vec.append(num.to_i)
  num = ''
  result = 1 if op == '*'

  if op != ''
    num_vec.each do |n|
      next unless n != 0

      case op
      when '+'
        result += n
      when '*'
        result *= n
      end
    end
    num_vec = []
  end

  sum_tot += result
  index -= 1
end

puts "The total sum of the problems is #{sum_tot}"
