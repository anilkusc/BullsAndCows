from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
import time
driver1 = webdriver.Chrome(".\chromedriver.exe")
driver1.get("http://localhost:8000")
time.sleep(1)
create = driver1.find_element_by_id("username")
create.send_keys(str(time.time()))
driver1.find_element_by_id("create").click()
time.sleep(1)
numberbar = driver1.find_element_by_id("numberbar")
numberbar.send_keys("1234")
driver1.find_element_by_id("readybutton").click()
session = driver1.find_element_by_id("session")
session = session.get_attribute('innerHTML')
driver2 = webdriver.Chrome(".\chromedriver.exe")
driver2.get("http://localhost:8000")
time.sleep(1)
joinusername = driver2.find_element_by_id("joinusername")
joinusername.send_keys(str(time.time()))
joinsession = driver2.find_element_by_id("joinsession")
joinsession.send_keys(session)
driver2.find_element_by_id("join").click()
time.sleep(1)
numberbar2 = driver2.find_element_by_id("numberbar")
numberbar2.send_keys("1234")
driver2.find_element_by_id("readybutton").click()
time.sleep(2)
predictionbar = driver1.find_element_by_id("predictionbar")
predictionbar.send_keys("4567")
driver1.find_element_by_id("submitbutton").click()
time.sleep(2)
predictionbar2 = driver2.find_element_by_id("predictionbar")
predictionbar2.send_keys("4567")
driver2.find_element_by_id("submitbutton").click()
time.sleep(2)
predictionbar3 = driver1.find_element_by_id("predictionbar")
predictionbar3.clear()
predictionbar3.send_keys("1234")
driver1.find_element_by_id("submitbutton").click()
#driver2.get("http://localhost:8000")
#driver1.close()