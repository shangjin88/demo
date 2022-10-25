

awk 'NR>2{print line}{line=$0} END{print line}'