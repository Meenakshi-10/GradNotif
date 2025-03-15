# 🎓 GradNotif - Reddit University Mention Notifier

This project is a Go-based Reddit scraper that monitors posts from the **"gradadmissions"** and **"mscs"** subreddits every 10 minutes to detect mentions of specific universities. The scraper runs concurrently for multiple subreddits, captures relevant posts containing target university names, and sends push notifications to your phone using Pushover. To automate execution, a Windows Task Scheduler job runs the compiled program (`GradNotif.exe`) at regular intervals, ensuring continuous monitoring without manual intervention.

---

## 🚀 Project Setup

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/Meenakshi-10/GradNotif.git
   cd GradNotif
   ```

2. **Install Go (if not already installed)**  
   Make sure you have Go installed on your system. You can download it from [Go Downloads](https://go.dev/dl/).  

3. **Install Dependencies**  
   The project uses the following libraries:  
   - [go-reddit](https://github.com/vartanbeno/go-reddit) for interacting with the Reddit API.  
   - [godotenv](https://github.com/joho/godotenv) to load environment variables from a `.env` file.  
   
   Run the following command to install the required libraries:  
   ```bash
   go mod init gradnotif
   go get github.com/vartanbeno/go-reddit/v2
   go get github.com/joho/godotenv
   ```

4. **Build the Project**  
   To build the executable, run:  
   ```bash
   go build -o GradNotif.exe main.go
   ```

5. **Run the Notifier**  
   Execute the binary to start scraping:  
   ```bash
   ./GradNotif.exe
   ```

6. **Automate the Notifier**  
   To run the notifier every 10 minutes, set up a Windows Task Scheduler job as follows:  
   - Open Task Scheduler and create a new basic task.  
   - Set the trigger to run every 10 minutes.  
   - Select **"Start a program"** and point to the compiled `GradNotif.exe` file.  
   - Make sure to check **"Run with highest privileges"** to avoid permission issues.  

---

## 🔔 Setting Up Pushover Notifications

To receive push notifications when a target university is mentioned, follow these steps:  

1. **Create a Pushover Account**  
   - Sign up at [Pushover.net](https://pushover.net/) and download the Pushover app on your mobile device.  

2. **Create a New Application**  
   - Go to your [Pushover Dashboard](https://pushover.net/apps) and click on **"Create an Application/API Token"**.  
   - Fill in the details and click **"Create Application"**.  
   - You will receive a **Pushover API Token**.  

3. **Obtain Your User Key**  
   - Your **User Key** can be found on your [Pushover Dashboard](https://pushover.net/).  

4. **Set Up Environment Variables**  
   - Create a `.env` file in the root of your project directory and add the following:  
     ```
     PUSHOVER_TOKEN=your_pushover_api_token
     PUSHOVER_USER_KEY=your_pushover_user_key
     TARGET_UNIVERSITIES=CMU,Stanford,UIUC,GaTech
     ```
   - Replace `your_pushover_api_token` and `your_pushover_user_key` with the actual values from your Pushover account.  
   - Update the `TARGET_UNIVERSITIES` list with the universities you want to monitor.  

---

## 📝 License
This project is licensed under the Apache License 2.0.

Feel free to reach out if you encounter any issues or have suggestions! Happy scraping! 😄

