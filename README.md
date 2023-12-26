<!-- Replace placeholders with your actual information -->
<h1>Go Image Compressor</h1>

<img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="GitHub license">

<h2>Description</h2>

<p>This Go program compresses images in a specified folder using the Tinify API. The compressed images are saved in a "Compressed" folder within the target directory.</p>

<h2>Features</h2>

<ul>
  <li>Uses the Tinify API for image compression.</li>
  <li>Simple command-line interface.</li>
  <li>Displays an ASCII art logo for a cool visual experience.</li>
</ul>

<h2>Installation</h2>

<ol>
  <li>Install Go on your machine.</li>
  <li>Clone this repository: <code>git clone https://github.com/sepehr-safaeian/Image-Compressor.git</code></li>
  <li>Change into the project directory: <code>cd Image-Compressor</code></li>
  <li>Run the program: <code>go run main.go</code></li>
</ol>

<h2>Configuration</h2>

<p>Before running the program, make sure to set your Tinify API Key in the <code>main.go</code> file:</p>

<pre>
<code>tinify.WithAPIKey("your-api-key-here")</code>
</pre>

<h2>Usage</h2>

<ol>
  <li>Enter the target folder address when prompted.</li>
  <li>The program will compress all images in the specified folder and save the compressed images in a "Compressed" folder.</li>
</ol>

<h2>Example</h2>

<pre>
<code>Location Folder Address: /path/to/images</code>
</pre>

<h2>Dependencies</h2>

<ul>
  <li><a href="https://github.com/moemoe89/go-helpers">moemoe89/go-helpers</a>: A package for image compression using the Tinify API.</li>
</ul>

<h2>License</h2>

<p>This project is licensed under the MIT License - see the <a href="LICENSE">LICENSE</a> file for details.</p>

