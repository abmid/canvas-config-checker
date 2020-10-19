<h1 align="center">Canvas Config Checker</h1>

<p align="center">
  <a href="https://travis-ci.com/abmid/canvas-config-checker.svg?branch=master"><img src="https://travis-ci.com/abmid/canvas-config-checker.svg?branch=master"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License"></a>
</p>
<p align="center"><img src="https://i.ibb.co/5FXCpzL/canvas-cc.png"></p>
A simple tool for check all configuration Canvas LMS before ready to build image

- This project beta version and part of UMM (University of Muhammadiyah Malang)

<h2>Installation / Use</h2>
<ol>
  <li>You must install Go lang first, check about installation in <a href="https://golang.org/doc/install" target="_blank">here</a>.</li>
  <li>Copy settings.yml.example, paste in same location and give name settings.yml</li>
  <li>Build app use command <code>make app</code> from your terminal, the result you can see in directory <code>dist</code>.</li>
  <li>Move to directory <code>dist</code> you will be see file <b>settings.yml</b> and <b>canvas-cc</b>, to run this app just write command <code>./canvas-cc</code> for *unix env</li>
</ol>
