#!/usr/bin/env ruby
# frozen_string_literal: true

require 'rspec/autorun'

EXAMPLE_INPUT = [1721, 979, 366, 299, 675, 1456].freeze

def solve_part1(input)
  input
    .combination(2)
    .find { |tuple| tuple.sum == 2020 }
    .inject(&:*)
end

def solve_part2(input)
  input
    .combination(3)
    .find { |tuple| tuple.sum == 2020 }
    .inject(&:*)
end

def main
  input = File.readlines(__dir__ + "/input.txt").map(&:to_i)

  puts solve_part1(input)
  puts solve_part2(input)
end

main

RSpec.describe 'day01' do
  it 'works for the first example' do
    result = solve_part1(EXAMPLE_INPUT)

    expect(result).to eq(514_579)
  end

  it 'works for the second example' do
    result = solve_part2(EXAMPLE_INPUT)

    expect(result).to eq(241_861_950)
  end
end
