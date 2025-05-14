import numpy as np

# Set parameters
mean = 100
std_dev = 15
num_samples = 1000000
output_file = "gaussian_integers.txt"

# Generate normally distributed values, then round them to integers
random_integers = np.random.normal(mean, std_dev, num_samples).astype(int)

# Save to file
with open(output_file, "w") as f:
    for number in random_integers:
        f.write(f"{number}\n")

print(f"Generated {num_samples} integer values with mean {mean} in {output_file}")
