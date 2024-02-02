
import json
import re

def summarize_json_structure_to_file(input_file_path, output_file_path):
    # Load JSON objects from the file
    with open(input_file_path, 'r') as file:
        json_objects = json.load(file)

    # Initialize a dictionary to store the summary of structures
    structure_summary = {}

    # Function to get the structure of a JSON object as a string
    def get_structure(obj):
        if isinstance(obj, dict):
            # Ignore 'metadata' key if present
            obj = {k: v for k, v in obj.items() if k != 'metadata'}
            
            # Sort the keys to ensure consistent structure representation
            keys = sorted(obj.keys())
            structure = "{" + ", ".join(f"{key}: {get_structure(obj[key])}" for key in keys) + "}"
            return structure
        elif isinstance(obj, list):
            # For lists, summarize the structure of the first element (assuming homogeneous lists)
            return "[" + (get_structure(obj[0]) if obj else "empty") + "]"
        else:
            # Instead of returning the type, just return 'value'
            return 'value'

    # Function to generate a REGEX pattern for a given structure
    def generate_regex_pattern(structure):
        # Replace placeholders with generic patterns
        pattern = structure.replace('value', '.+?')
        pattern = pattern.replace(' ', '\s*') # Allow for any whitespace
        return re.compile(pattern)

    # Analyze each JSON object and count the occurrences of each structure
    for obj in json_objects:
        structure = get_structure(obj)
        if structure not in structure_summary:
            structure_summary[structure] = {'count': 0, 'regex': generate_regex_pattern(structure)}
        structure_summary[structure]['count'] += 1

    # Sort the summary by count in descending order
    sorted_summary = dict(sorted(structure_summary.items(), key=lambda item: item[1]['count'], reverse=True))

    # Write the sorted summary to an output file
    with open(output_file_path, 'w') as file:
        # Convert regex objects to string before dumping to JSON
        for structure in sorted_summary:
            sorted_summary[structure]['regex'] = sorted_summary[structure]['regex'].pattern
        json.dump(sorted_summary, file, indent=4)

    return output_file_path

# Example usage
if __name__ == "__main__":
    input_file_path = 'Kustomer3.json'  # Replace with your input file path
    output_file_path = 'sorted_summary_with_regex_of_json_structures.json'

    # Run the function
    output_file_path = summarize_json_structure_to_file(input_file_path, output_file_path)
    print(f"Summary saved to {output_file_path}")
