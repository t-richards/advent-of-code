#!/usr/bin/env ruby

require 'rspec/autorun'
require 'pry'

EXAMPLE_INPUT = [
  '1-3 a: abcde',
  '1-3 b: cdefg',
  '2-9 c: ccccccccc'
].freeze

POLICY_SPLITTER = Regexp.union(': ', '-', ':', ' ')

def solve_part1(input)
  input.count do |line|
    min, max, char, password = line.split(POLICY_SPLITTER)
    min = min.to_i
    max = max.to_i
    count = password.count(char[0])
    min <= count && count <= max
  end
end

def solve_part2(input)
  input.count do |line|
    min, max, char, password = line.split(POLICY_SPLITTER)
    min = min.to_i
    max = max.to_i
    [min, max].one? { |i| password[i - 1] == char }
  end
end

def main
  input = File.readlines("#{__dir__}/input.txt").map(&:strip)

  puts solve_part1(input)
  puts solve_part2(input)
end

main

RSpec.describe 'day02' do
  it 'works for the first example' do
    result = solve_part1(EXAMPLE_INPUT)

    expect(result).to eq(2)
  end

  it 'works for the second example' do
    result = solve_part2(EXAMPLE_INPUT)

    expect(result).to eq(1)
  end
end
