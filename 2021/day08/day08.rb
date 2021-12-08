#!/usr/bin/env ruby

require 'rspec/autorun'
require 'pry'

EXAMPLE_INPUT = [
  'be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe',
  'edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc',
  'fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg',
  'fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb',
  'aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea',
  'fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb',
  'dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe',
  'bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef',
  'egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb',
  'gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce'
].freeze

LEN_2_NUM = {
  2 => 1,
  3 => 7,
  4 => 4,
  7 => 8
}.freeze

P1_LENS = [2, 3, 4, 7].freeze

def solve_part1(input)
  input.sum do |line|
    line
      .split('|')
      .last
      .split
      .count { |w| P1_LENS.include?(w.length) }
  end
end

def solve_part2(input)
  input.sum do |line|
    input, output = line.split('|').map(&:split)
    easy_nums = input.map { |w| LEN_2_NUM[w.length] }
    mapping = input.zip(easy_nums).to_h

    # Now figure out the 5-length patterns...
    binding.pry
  end
end

def main
  input = File.readlines("#{__dir__}/input.txt")

  puts solve_part1(input)
  puts solve_part2(input)
end

# main

RSpec.describe 'day08' do
  it 'works for the first example' do
    result = solve_part1(EXAMPLE_INPUT)

    expect(result).to eq(26)
  end

  it 'works for the second example' do
    result = solve_part2(EXAMPLE_INPUT)

    expect(result).to eq(5353)
  end
end
