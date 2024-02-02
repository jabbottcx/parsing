"""
This module provides functionality to analyze the structure of JSON objects within a given JSON file.
It summarizes the unique structures found, counts the occurrences of each structure, generates REGEX patterns
to match objects with those structures, and saves this summary to an output file.

Usage:
    python json_analysis_script_commented.py <input_json_file.json> <output_summary_file.json>
"""

import json
import re
import sys

def summarize_json_structure_to_file(input_file_path, output_file_path):
    """
    Analyzes the structures of JSON objects in the specified input file, counts the occurrences of each unique structure,
    generates REGEX patterns for identifying matching objects, and writes this summary to the specified output file.

    Parameters:
    - input_file_path (str): Path to the input JSON file.
    - output_file_path (str): Path for the output file to save the summary.

    The function reads the input JSON file, processes each object to determine its structure, and then
    compiles a summary that includes the count of occurrences and a REGEX pattern for each unique structure.
    This summary is then saved to the output file in JSON format.
    """
    # Load JSON objects from the input file
    with open(input_file_path, 'r') as file:
        json_objects = json.load(file)

    # Dictionary to store the summary of JSON object structures
    structure_summary = {}

    def get_structure(obj):
        """
        Determines the structure of a JSON object and represents it as a string.

        This function recursively analyzes a JSON object (which can be a dictionary, a list, or a basic data type),
        constructing a string that represents its structure. The structure of dictionaries includes the keys and
        recursively determined structures of their values. Lists are assumed to be homogeneous, and their structure
        is represented based on the first element.

        Parameters:
        - obj: The JSON object to analyze.

        Returns:
        - A string representation of the JSON object's structure.
        """
        if isinstance(obj, dict):
            # Exclude 'metadata' key from analysis if present
            obj = {k: v for k, v in obj.items() if k != 'metadata'}
            
            # Construct structure string for dictionary objects
            keys = sorted(obj.keys())
            structure = "{" + ", ".join(f"{key}: {get_structure(obj[key])}" for key in keys) + "}"
            return structure
        elif isinstance(obj, list):
            # Handle list objects by analyzing the structure of the first element
            return "[" + (get_structure(obj[0]) if obj else "empty") + "]"
        else:
            # Basic data types are generalized as 'value'
            return 'value'

    def generate_regex_pattern(structure):
        """
        Generates a REGEX pattern from a structure string to facilitate matching similar structures.

        This function takes a string representation of a JSON object's structure and creates a REGEX pattern
        that can be used to identify JSON objects with a matching structure in the text format.

        Parameters:
        - structure (str): A string representing a JSON object's structure.

        Returns:
        - A compiled REGEX pattern object.
        """
        # Replace placeholders with generic patterns, considering optional whitespace
        pattern = re.compile(structure.replace('value', '.+?').replace(' ', '\\s*'))
        return pattern

    # Analyze and summarize the structures of all JSON objects
    for obj in json_objects:
        structure = get_structure(obj)
        if structure not in structure_summary:
            structure_summary[structure] = {'count': 0, 'regex': generate_regex_pattern(structure)}
        structure_summary[structure]['count'] += 1

    # Sort the summary by count in descending order and prepare for output
    sorted_summary = {k: v for k, v in sorted(structure_summary.items(), key=lambda item: item[1]['count'], reverse=True)}
    for structure in sorted_summary:
        sorted_summary[structure]['regex'] = sorted_summary[structure]['regex'].pattern

    # Write the structured summary to the output file
    with open(output_file_path, 'w') as file:
        json.dump(sorted_summary, file, indent=4)

if __name__ == "__main__":
    # Command-line usage validation
    if len(sys.argv) < 3:
        print("Usage: python json_analysis_script_commented.py <input_json_file.json> <output_summary_file.json>")
        sys.exit(1)

    # Execute the main function with command-line arguments
    summarize_json_structure_to_file(sys.argv[1], sys.argv[2])
    print(f"Summary saved to {sys.argv[2]}")
