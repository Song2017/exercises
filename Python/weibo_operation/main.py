import pdb

from selenium import webdriver
import time
from selenium.webdriver.common.by import By

browser = webdriver.Chrome()


# 登录微博
def weibo_login(username_, password_):
    # 打开微博登录页
    browser.get('https://passport.weibo.cn/signin/login')
    browser.implicitly_wait(5)
    time.sleep(1)
    # 填写登录信息：用户名、密码
    browser.find_element(value="loginName").send_keys(username_)
    browser.find_element(value="loginPassword").send_keys(password_)
    time.sleep(1)

    # 点击登录
    browser.find_element(value="loginAction").click()
    time.sleep(1)


# 添加指定的用户
def add_follow(uid):
    browser.get('https://m.weibo.com/u/' + str(uid))
    time.sleep(1)
    # browser.find_element_by_id("follow").click()
    follow_button = browser.find_element(
        value='//div[@class="m-add-box m-followBtn"]')
    follow_button.click()
    time.sleep(1)
    # 选择分组
    group_button = browser.find_element(
        value='//div[@class="m-btn m-btn-white m-btn-text-black"]')
    group_button.click()
    time.sleep(1)


# 给指定某条微博添加内容
def add_comment(weibo_url, content):
    browser.get(weibo_url)
    browser.implicitly_wait(5)
    browser.find_element(
        by=By.CSS_SELECTOR, value="textarea.W_input").clear()
    browser.find_element(
        by=By.CSS_SELECTOR, value="textarea.W_input").send_keys(
        content)
    time.sleep(2)
    browser.find_element(
        by=By.CSS_SELECTOR, value=".W_btn_a").click()
    time.sleep(1)


# 发文字微博
def post_weibo(content_):
    # 跳转到用户的首页
    browser.get('https://weibo.com')
    browser.implicitly_wait(5)
    # 点击右上角的发布按钮
    browser.find_element(
        by=By.CSS_SELECTOR, value="[node-type='publish']").click()
    # 在弹出的文本框中输入内容
    browser.find_element(
        by=By.CSS_SELECTOR, value="textarea.W_input").send_keys(content_)
    time.sleep(2)
    # 点击发布按钮
    browser.find_element(
        by=By.CSS_SELECTOR, value="[node-type='submit']").click()
    time.sleep(1)


if __name__ == '__main__':
    username = '18765918310'
    password = "son过*"
    weibo_login(username, password)
    pdb.set_trace()

    # 每天学点心理学UID
    uid = '1890826225'
    # add_follow(uid)

    # 给指定的微博写评论
    weibo_url = 'https://weibo.com/1890826225/HjjqSahwl'
    content = 'Gook Luck!好运已上路！'
    # 自动发微博
    content = '每天学点心理学'
    # post_weibo(content)
