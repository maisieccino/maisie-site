---
id: how-i-work-obsidian
aliases:
  - "How I Work: Obsidian"
  - "How I work: Obsidian"
tags:
  - technology
  - neovim
published: false
title: "How I work: Obsidian"
---
# How I Work: Obsidian
I'm a disorganised mess.

It seems a bit weird to start a blog post off by insulting myself, but it's true! I think it's not really a bad thing. Noticing and accepting it is really useful, because I can find some comfort in knowing that whatever organisational system or notekeeping I use, it's not going to change.

I'm a disorganised mess and that's okay. I don't have to keep trying to search for the "perfect" system.

Anyway, let me tell you about how I use Obsidian to cram together all of my thoughts, ideas, and things I learn. It's messy and chaotic. Importantly though, it's _mine_.

## What is Obsidian?
[Obsidian](https://obsidian.md) is a free (but closed-source) program for notes. It essentially gives you folders full of text files that you can write thoughts into. Your notes are written in a special format called _Markdown_, which lets you add things like headings, bold/italic/strikethrough text, lists, that kind of thing. The special sauce of Obsidian is how you can use tags and links to start building connections between notes. So instead of just a bunch of random text files, you can quickly end up with an entire knowledge base of ideas and things you learned, and use tags and folders to keep similar things together.

This openness to how the system works means you have a lot of creative freedom in how you organise your collection (they call it a "vault"). A lot of creative freedom... to make absolutely horrendous choices in your structure.

A bunch of people have come forward with their ideas of how the "perfect second brain" should be structured and organised. You get things like the "PARA method", "Zettelkasten", "Map of contents", and so many other methods that have cool but nerdy-sounding names. And, like any decision that has opinions, everyone likes to politely discuss* why their method is the best.

I have kinda chosen to not follow that.

## My organisational system
Ha.

I have a loose system with a numbered set of top-level folders:

```
/notes
├── 050 Mind
│   └── 051 Soothing Skills
├── 100 Projects
│   ├── Maisie site
│   └── youtube
├── 200 Inbox
├── 400 Ideas
├── 500 Knowledge
│   ├── Command line stuff
│   ├── Health
│   ├── iOS
│   ├── Linux
│   ├── Neovim
│   ├── Obsidian
│   └── Web
├── 600 Daily Notes
├── 700 Blog
├── 800 Clippings
├── 900 Places
├── Assets
├── Templates
```

I try to keep things roughly organised by topic. `Ideas` is the folder for things I want to do at some point. `Projects` for things I'm actually working on. `Knowledge` for the majority of the things I learn along the way. `Mind` is a folder where I keep all of my thoughts on stuff like therapy, who I am, what do I want to be. `Daily notes` is a daily tracker of what I'm up to on a given day and is quite messy. `Inbox` is pretty cool but we'll come back to that later.

## Cross-device syncing
I read and write to my notes from four places--my phone, my tablet, my laptop and my desktop computer. I want to be able to sync my notes between them in the easiest way.

Obsidian supports syncing via iCloud for free, but this would stop me from being able to sync with my Linux desktop. There's a hosted sync service offered by Obsidian, but it involves paying yet another subscription fee, so I'm not yet keen on that. And using a third party cloud storage system can be a bit dangerous! People have shared cautionary tales of their files disappearing into the ether and not being able to get them back.

I decided to use Git.

Since Obsidian is just a pile of text files, I thought why not? Using Git also means that I have version control for free, it works on almost any device, and I have data redundancy. If my computer caught fire, I'd still have all of my notes safely stored on my other devices. This seems like the best way forward for what I want right now.

### Desktop setup
Once I had my local vault created in Obsidian, I needed to set up Git.

```sh
git init
```

Oh huh okay. That was easy.

I set up a quite little script to quickly commit and push my notes, including the date in the message.

```bash
alias syncnotes="git add . && git commit -m \"Sync mac $(date -I)\" && git push"
```

If I was feeling especially fancy, I could write a cron job to do this at regular intervals, but I haven't needed to do that yet.

### iOS
Making Obsidian work with Git on mobile is a bit more fiddly. The Obsidian mobile app _does_ have a Git plugin! But it seemed to have a tendency to start deleting my files randomly and without notice, before committing them to my Git repository. Not ideal.

I ended up paying for Working Copy, a Git app for Apple devices. Firstly, I created a new blank vault in the Obsidian app. Then I opened Working Copy, and pointed it at my new folder. I set up a SSH key so it could write to my remote Git repository, and then reset it to the latest commit from the remote. I opened the Obsidian app and there were all my notes!

I now needed a way to automate syncing my notes with the remote repository. Working Copy supports the Shortcuts app, which is exactly what I needed. I have one shortcut that runs whenever I open the Obsidian app, that tells Working Copy to pull down the latest changes. I have another one I use that will commit any local changes and push them up. Really straightforward and seems to do the trick nicely.

![[Pasted image 20250914111348.png]]

## Plugins
Obsidian has an extremely strong ecosystem of plugins. For pretty much anything you'd need concerning notes, someone's probably created a plugin for it. I don't use many, but I've included a couple that I find really handy.

### Calendar
![[Pasted image 20250914111704.png]]

The calendar plugin presents all of my daily notes as a calendar to let me navigate through them nice and quickly. The dots signify to-do list items: filled in means completed tasks, empty dots mean there's incomplete tasks. Straightforward and does the task nicely.

### Dataview
Dataview is a plugin that lets you write queries for your vault. Obsidian's since released its [Bases feature](https://help.obsidian.md/bases), which solves a lot of problems Dataview tried to solve. But Dataview is still pretty neat for a few reasons.

![[Pasted image 20250914112052.png]]

You can write queries that perform maths and other cool stuff on your notes and properties. If you've ever written SQL for querying a database, it's really similar!

The query above for tracking my sleep looks like this:

```dataview
list without id
  "Average sleep, last 7 days: " + sum(rows.sleep_duration) / length(rows.sleep_duration)
from "600 Daily Notes"
where file.ctime >= date(today) - dur(7d)
and row.sleep_duration
group by ""
```

The nerd in me loves this.

## Writing notes: Neovim
![[Pasted image 20250914112453.png]]

As I mentioned earlier, my favourite thing about Obsidian is that it's a bunch of Markdown files. This means that you can use your favourite text editor to write! So naturally, I set up an Obsidian editing system directly in Neovim.

[Obsidian.nvim](https://github.com/obsidian-nvim/obsidian.nvim) is a plugin that adds a lot of really nice functionality for navigating your vault. Each file has a footer that shows properties and backlinks. It has settings that let you automatically generate properties for a note. You can search through all your tags and tagged notes using your favourite file picker. There's autocomplete for adding links to other files. There's shortcuts for jumping to your daily note. You can even paste images into your vault straight from the clipboard!

It's not perfect for sure, but writing notes in Neovim is really fun and makes me _want_ to actually write.

I use [render-markdown.nvim](https://github.com/MeanderingProgrammer/render-markdown.nvim) as the main way to format and render markdown in the editor. It will make bullet points, headings and even hyperlinks look a lot more tidy. Code blocks get the syntax highlighting for the language you specify.

![[Pasted image 20250914113302.png]]

[snacks.nvim](https://github.com/folke/snacks.nvim) is a huge plugin with so many different features. I use it as my file picker and explorer mainly. It also comes with image support. This means that if I include an image in my note, it'll show up embedded in the terminal. Which is absolutely delightful.

Finally, I use [conform.nvim](https://github.com/stevearc/conform.nvim) to let me format my Markdown to keep it consistent, `Markdownlint` to alert me of any problems, and [no-neck-pain](https://github.com/shortcuts/no-neck-pain.nvim) aligns the buffer to them middle of the window so that it's a bit nicer for writing longer paragraphs and files.

Here's the key shortcuts I've set up for Obsidian:

```lua
    keys = {
      { "<localleader>ot", "<cmd>Obsidian today<CR>", desc = "Daily note (today)" },
      { "<localleader>o+", "<cmd>Obsidian tomorrow<CR>", desc = "Daily note (tomorrow)" },
      { "<localleader>o-", "<cmd>Obsidian yesterday<CR>", desc = "Daily note (yesterday)" },
      { "<localleader>on", "<cmd>Obsidian new<CR>", desc = "New" },
      { "<localleader>or", "<cmd>Obsidian rename<CR>", desc = "Rename" },
      { "<localleader>oP", "<cmd>Obsidian paste_img<CR>", desc = "Paste image" },
      { "<localleader>oT", "<cmd>Obsidian tags<CR>", desc = "Tags" },
      { "<localleader>om", "<cmd>Obsidian template<CR>", desc = "Insert template" },
      { "<localleader>ob", "<cmd>Obsidian backlinks<CR>", desc = "Backlinks" },
      { "<localleader>oc", "<cmd>Obsidian toc<CR>", desc = "Insert ToC" },
      { "<localleader>ox", "<cmd>Obsidian extract_note<CR>", desc = "Extract note", mode = "v" },
      { "<localleader>o ", "<cmd>!/bin/zsh -c syncnotes<CR>", desc = "Push notes to remote" },
    },
```

My full configuration is available [on GitHub](https://github.com/maisieccino/hello-computer/blob/master/vim/.config/nvim/lua/plugins/obsidian.lua).

## Shortcuts
I use a small set of shortcuts to add some extra smart to my Vault.

Firstly, newer iPhones come with a programmable "action button" on the side of the device. I've set this to open the url `obsidian://daily`, which jumps straight to today's note. This makes it really easy to jot things down and check on my to-do list as I'm out and about!

The "update today's note" runs in the morning and the evening. It will search my Apple Health data for how much I slept the night before, as well as how much water and caffeine I've consumed, and adds them to my daily note as properties. This is fantastic for me, as I can see this at a glance, and if I haven't slept well, I can be reminded why I might be feeling a little bit cranky and to go easy on myself. At some point I've love to make some visualisations with this data, but that day is not today.

"My newsletters" runs first thing in the morning too. I've recorded a list of RSS feeds I want to follow. When the shortcut runs, it will go through each feed, filter to any article written in the previous day, before converting it to Markdown and saving it into my vault, in the `Inbox` folder. I filter through using the `read: false` property to give myself a cheap, janky feed reader that I can read anywhere!

![[Pasted image 20250914132817.png]]

Finally, I'm building smaller shortcuts which should help me with summarising what I've been up to in a week, shortcuts to capture links, and also ways to store places I've visited and ideas I have. Watch this space.

I've always believed Shortcuts to be one of the coolest things you can play around with on iPhones. More and more app developers seem to be starting to build support for tasks, and it means you can start automating and sticking together a lot of useful functions and utilities. Highly underrated.

---

This should give you a rough overview of how I use Obsidian to try and collect what I've learned and what I'm working on. I think it's a really fantastic tool! But it's also so easy to overthink your setup. Starting small, and focusing on just writing and getting ideas down, is the best thing you can do. Keep it open and active near to you as much as you can, and you quickly find it starts building up. I've tried changing my structure a few times, and it never really seems too forced or painful.

And it's really important to remember that these are all text files at the end of the day. That means you can, if you want, use all of your favourite command line tools to navigate around. `fzf` for fuzzy finding a note, `grep` (and friends) for searching text, you can even use `yq` to filter through the properties on your notes:

```bash
 » yq --front-matter=extract '.tags' 500\ Knowledge/Neovim/glance.md
- neovim
- plugin
```

There's so many possibilities here and I hope you have a lot of fun customising your vault.
