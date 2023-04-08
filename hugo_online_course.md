In this example, we start by defining some metadata for the course using Hugo's front matter syntax. We specify the course title, description, 
duration, and instructor.

We then define two sections for the template: one for the course details and one for the course content. In the course details section, we use 
Bootstrap classes to create a responsive layout with two columns. We output the course title, description, duration, and instructor using Hugo's dot 
notation to access the metadata we defined earlier.

In the course content section, we use a range loop to iterate over the content elements for the course. We check the type of each element (chapter, 
section, or video) and output the appropriate HTML markup. For videos, we use an iframe to embed the video using the URL specified in the content 
element.

This is just a basic example, and you can customize it to fit your specific needs. You can use Hugo's built-in shortcodes and other template features 
to add more complex functionality to your online course template.


```
---
title: "Introduction to Computer Science"
description: "Learn the fundamentals of computer science with our comprehensive online course."
course_duration: "8 weeks"
course_instructor: "John Doe"
---
{{- $course := . -}}

<section class="course-details">
  <div class="container">
    <div class="row">
      <div class="col-md-6">
        <h1>{{ $course.Title }}</h1>
        <p>{{ $course.Description }}</p>
        <p><strong>Duration:</strong> {{ $course.Course_Duration }}</p>
        <p><strong>Instructor:</strong> {{ $course.Course_Instructor }}</p>
      </div>
    </div>
  </div>
</section>

<section class="course-content">
  <div class="container">
    <div class="row">
      <div class="col-md-8">
        {{- range $index, $element := .Content -}}
          {{- if eq $element.Type "chapter" -}}
            <h2>{{ $element.Title }}</h2>
          {{- else if eq $element.Type "section" -}}
            <h3>{{ $element.Title }}</h3>
          {{- else if eq $element.Type "video" -}}
            <h4>{{ $element.Title }}</h4>
            <div class="embed-responsive embed-responsive-16by9">
              <iframe class="embed-responsive-item" src="{{ $element.Video_URL }}"></iframe>
            </div>
          {{- end -}}
        {{- end -}}
      </div>
    </div>
  </div>
</section>
```
