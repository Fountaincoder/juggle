import random

# Set parameters
p_heads = 0.7  # Probability of getting heads
num_spins = 1000000
output_file = "bent_coin_spins1000000.txt"

# Generate spins
with open(output_file, "w") as f:
    for _ in range(num_spins):
        flip = "1" if random.random() > p_heads else "0"
        f.write(f"{flip}\n")

print(f"Generated {num_spins} bent coin spins in {output_file}")
