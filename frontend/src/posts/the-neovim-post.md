---
id: Neovim post
aliases:
  - The Neovim post
tags:
  - tech
authors:
  - Maisie Bell
feature_image: https://blog.mbell.dev/content/images/2025/10/pasted_image_20251024165336.png
hash: 389dc4743c05ad3464e6d9b4340fca47c3102e7f9e0c3d31810b2102322fa916
post_id: 68fcdebd88baba00019ef402
status: published
tiers:
  - Free
  - Default Product
title: The Neovim post
---

I think there's something really satisfying and enjoyable about taking a tool you use every day, and customising it and tuning it so it becomes truly your own. Knowing that all of the features and abilities are just what you need and nothing more. Weird little parts that fit your own quirks. Style that matches how you feel.

This is the blog post about Neovim.

![A Neovim window with a "popup" window showing a tool called LazyGit](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251024165421.png)
A Neovim window with a "popup" window showing a tool called LazyGit

## What even is Neovim?
[Neovim](https://neovim.io) is a program that lets you edit text. It's based on Vim, a program that has been around for years, which itself is based on Vi. All these programs run from the command line and draw text on the screen, instead of running in its own window. Usually, you do this by opening a _terminal_ and running `nvim`. It will then launch and take over the entire window with its interface.

Vi, vim and neovim have kinda built a cult following as the way you move around documents, and select and edit text is just _so different_ from anything else. Either you love it or you hate it. There's three "modes":
* "normal" for moving your cursor around the text file
* "insert" for adding text
* "visual" for selecting text

You move the cursor around with `h` for left, `j` for down, `k` for up and `l` for down. Why not just the arrow keys I hear you ask? It's because `hjkl` sit on the home row, where your fingers naturally rest during touch typing. So you don't have to move your fingers off the home row to start moving around. You can then type a number before a key to move that number of spaces. `30j` moves you down 30 lines.

Chaining key presses like this is known as "vim motions" and is one of the core features that makes Vim so different from anything else. Jump to the start of your line with `0` and to the end with `$`. Top of the document is `gg`. Start of the next word with `w` and the end of the word with `e`. `f<key>` lets you jump to the next `<key>` in the file. `10f(` jumps to the 10th "`(`" after your cursor. It's surprisingly intuitive and you start learning you can connect things together and build really graceful movements. If it works for you, it makes navigating and editing text files really easy!

Neovim has a few extra useful bits with it over regular Vim. Configuration files can be written using a language called `Lua`, which has made it easier than ever for people to start building really neat plugins and features. It has built in ways of communicating with Python programs to extend its usefulness further. It has support built in for connecting to "language servers", tools that let you get code completion, syntax error messages, and one-tap formatting.

IDK if this sounds like fervent fangirling or if it's actually selling you on Neovim. Either way, I'll crack on.

## How I'm using Neovim
Neovim is the program I use to write stuff when I'm on my laptop. I've found other programs will usually come with a plugin to provide vim-style keyboard shortcuts, but they often end up lacking _something_. Call it a _je ne sais quoi_ if you must.

Besides, using the same program for pretty much all my writing needs means that I've become incredibly familiar with the controls. Everything is customised to my life and how I use my computer. As a result there's going to be stuff I use or do that will be bewildering to anyone else. I've also set up certain features that will only work when I'm at work or when I'm at home. This means that I'm able to code at work in a familiar, fast interface. And at home, I'm not distracted with tools that I wouldn't otherwise use.

My configuration files are available on [GitHub](https://github.com/maisieccino/hello-computer/tree/master/vim/.config/nvim).

### The dashboard
![My Neovim dashboard screen. Showing a chart, shortcuts and a pixellated image](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025141217.png)
My Neovim dashboard screen. Showing a chart, shortcuts and a pixellated image

Running `nvim` brings up my home dashboard. This is powered by a plugin called [Snacks.nvim](https://github.com/folke/snacks.nvim/), which provides a whole bunch of pickers, file explorers, and even image support. I use a program called `chafa` to render a photo I took in Spain as text, just because it looks pretty. A list of actions lets me jump straight into parts of the program I use often. I then built a small Go program, which reads in my daily journal notes, grabs how much water and exercise I've done, then prints it out as a pretty chart. This helps me stay on track with my habits! The library I use is called [ntcharts](https://github.com/NimbleMarkets/ntcharts), for those that want to build their own.

Finally, a little display at the bottom lists any GitHub notifications I have. As I'm trying to do a little bit more open source work, this is really useful to stay on track with conversations and pull requests!

### Coding
![A split window showing CSS code, javscript code as well as a search window.](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025142344.png)
A split window showing CSS code, javscript code as well as a search window.

Editing code in Neovim feels really comfortable for me! It already works pretty well out of the box, and then there's a really good ecosystem of plugins that make it even better. Neovim supports the _Language Server Protocol_, a standard way for text editors to talk with servers that provide information about the code being edited. This includes formatting, finding symbols (variables and function names) and where things are defined and used. TreeSitter is a library built into Neovim that lets it build _Abstract Syntax Trees_. This means that the library is able to take a piece of code, parse it and understand hierarchy, symbols and the general flow of the code. The "tree" it creates is this big list of all the different symbols and any sub-symbols. This lets Neovim set syntax highlighting, and what the code actually _means_, so that you can easily jump around it with your fancy Vim key movements. The image below shows some Typescript code, with the syntax tree printed on the right hand side.

![A screenshot of Neovim with the syntax tree printed on the right hand side.](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025143042.png)
A screenshot of Neovim with the syntax tree printed on the right hand side.

A few plugins I also find useful. I can surround text with brackets, quotes, etc just by typing `sa[`. I have a shortcut set to show me the "outline" of my code, a list of all the functions and classes and variables so I can easily navigate a file. `[space] + v` will comment out a line. `blink.cmp` adds a really fast and handy popup window for autocompleting my text. There's even a tool called `tabs-vs-spaces.nvim` so that you don't have to care about mixing the two up or using the wrong one - it will just fix it. Finally, `glance.nvim` lets me "peek" at the definition of a variable/function, or where it's being used, without losing my place in the file I'm editing.

![A screenshot showing a glance pop up, displaying a list of references to a function](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025143703.png)
A screenshot showing a glance pop up, displaying a list of references to a function

### Note-taking
In my [Obsidian](https://mbell.dev/post/how-i-work-obsidian/) post, I talked about how I'm using this program to manage my notes as a big folder of text files. I actually end up using Neovim to do most of my typing and notetaking! Because it is just a bunch of text files at the end of the day, it means that it wasn't a lot of work to get this working in Neovim. I use a plugin called `obsidian.nvim` to help with linking between notes, collecting tags, and working with templates. It will even let me paste images into my note vault, straight from my clipboard!

Using `render-markdown.nvim`, my text is set out a bit nicer. Callouts, tables and bullet points/checkboxes are all pretty and aligned. `no-neck-pain.nvim` crops the editor window to a nice width for writing and reading text. `Snacks.nvim` will render images in the terminal, and even take LaTeX and render it! The result is a really delightful note-taking experience which is fast and straightforward. It makes me actually _want_ to write stuff down.

![Editing notes (and this blog post!) in Neovim.](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025144045.png)
Editing notes (and this blog post!) in Neovim.

### Working with GitHub
`Octo.nvim` is a plugin I've really enjoyed using lately. It allows you to fetch pull requests and issues from GitHub and interact with them, straight from Neovim! You can take a PR, and even review it line by line in your editor. It's honestly incredible.

![A PR conversation in Octo.](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025144900.png)
A PR conversation in Octo.

![Looking at a PR change](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025144716.png)
Looking at a PR change

Honestly, finding this and using it blew my mind.

### Jupyter notebooks
Finally, maybe the coolest thing I've been able to get working in Neovim. By combining a set of plugins--`quarto`, `molten`, and `jupytext`--I've been able to get a working setup for reading, editing and running entire Jupyter notebooks, as if they were just markdown files. [This](https://github.com/benlubas/molten-nvim/blob/main/docs/Notebook-Setup.md) guide explains the process well.

![A screenshot of a graph generated from a Jupyter notebooks](https://blog.mbell.dev/content/images/2025/10/pasted_image_20251025145728.png)
A screenshot of a graph generated from a Jupyter notebooks

The experience is a little bit buggy but it's honestly so cool.

## Now, your turn
Do I think that everyone reading this post should install Neovim and try it out? I mean sure! Trying new things is pretty fun and you might find it really enjoyable. But it's not the main message I want to share.

Whatever tools you use in your personal, creative or work endeavours, I think there's a lot of value in experimenting and tweaking and making it completely _yours_. I think it feels really rewarding knowing that your tool works the best way possible for yourself. Even just messing around with the home screen on your phone can be really fun! Recent iPhones come with an "action button" which you can program to run any shortcut you like. What's something you do often that you could automate? Maybe you can press the button to record a "note to self" which is then transcribed to text and saved to a note for later. Top marks if you talk like Dale Cooper from Twin Peaks.

Our digital lives are absolutely inundated with generative AI and LLMs. Isn't it about time that we start making computing more fun and personal again? While we're at it, maybe go build a fun little NeoCities website or a blog or something and share it with me. I'll be delighted.
