#!/usr/bin/python
# Filename: getmm.py

import urllib.request
import os
import re
import time

myurl = input("please input the url you want: ")
# mypage = int(input("how many pics you want:"))

def url_open(url):
	req = urllib.request.Request(url)
	req.add_header('User-Agent','Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6')
	response = urllib.request.urlopen(url)
	html = response.read()

	return html

def get_page(url):
    html = url_open(url).decode('ISO-8859-1')
    b = re.search(r'(\d+_\d{1,3}).html',html)
    return b.group(1)


def find_img(url):
	html = url_open(url).decode('ISO-8859-1')
	img_addrs = []
	a = re.search(r'<img alt=.*src="(\S+.jpg)"',html)
	img_addrs.append(a.group(1))
	return img_addrs

	

def save_imgs(folder, img_addres):
    for each in img_addres:
        print (each)
        filename = each.split('/')[-1]
        with open(filename,'wb') as f:
            img = url_open(each)
            f.write(img)


def download_mm(folder='MM',url=myurl):
    # pages =pages +1
    if os.path.exists(folder):
        pass
    else:
        os.mkdir(folder)
            
    os.chdir(folder)
    judge =[]
    while True:
        nextid = get_page(url)
        # get raw url
        aa = re.search(r'(\S+/)(\d+/)\d+',url)# 2018/26825
        raw_url = aa.group(1)
        year = aa.group(2)


        # page_url = raw_url + page_num +'.html'
        #print (page_url)
        img_addres = find_img(url)
        save_imgs(folder, img_addres)
        if nextid in judge:
            break
        else:
            judge.append(nextid)
            url = raw_url + year + nextid +'.html'
        #time.sleep(5)
        

if __name__ == '__main__':
	download_mm()