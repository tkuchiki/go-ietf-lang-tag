#!/usr/bin/env ruby

require "csv"

data = {}
CSV.foreach(ARGV[0]) do |row|
  next if row[0] == "lang"
  data[row[1]] = [] if data[row[1]].nil?
  data[row[1]].push(row[0])
end

require "pp"
data.each do |lang_type, langs|
  quoted_langs = langs.map { |s| "\"#{s}\""}.join(',')
  puts "\"#{lang_type}\": []string\{#{quoted_langs}\},"
end
