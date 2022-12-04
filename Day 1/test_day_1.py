from main import *

def test_parse_cals_per_elf():
    test_input = '''
      1000
      2000
      3000

      4000

      5000
      6000

      7000
      8000
      9000

      10000
    '''

    assert parse_cals_per_elf(test_input) == [6000, 4000, 11000, 24000, 10000]

def test_max_total_cals():
  cals_per_elf = [6000, 4000, 11000, 24000, 10000]
  assert max_cals(cals_per_elf) == 24000

def test_total_top_n_cals():
  cals_per_elf = [6000, 4000, 11000, 24000, 10000]
  assert total_top_n_cals(cals_per_elf, 3) == 45000
