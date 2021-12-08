#!/usr/bin/env ruby

require 'rspec/autorun'
require 'pry'

EXAMPLE_INPUT = [3, 4, 3, 1, 2].freeze

def lanternfish(input, days)
  buckets = [0] * 9
  input.map { |i| buckets[i] += 1 }

  days.times do
    buckets.rotate!
    buckets[6] += buckets[8]
  end

  buckets.sum
end

def solve_part1(input)
  lanternfish(input, 80)
end

def solve_part2(input)
  lanternfish(input, 256)
end

def main
  input = File.read("#{__dir__}/input.txt").split(',').map(&:to_i)

  puts solve_part1(input)
  puts solve_part2(input)
end

main

RSpec.describe 'day06' do
  it 'works for the first example' do
    result = solve_part1(EXAMPLE_INPUT)

    expect(result).to eq(5934)
  end

  it 'works for the second example' do
    result = solve_part2(EXAMPLE_INPUT)

    expect(result).to eq(26_984_457_539)
  end
end
