# squeakywheel
checks auditd rules by firing each off once

# perms_checker

This Go binary performs the following checks:

Dummy File Access: Attempts to read and write to a specified dummy file (default: /root/ring0_rights/user_group/dummy_file). This tests for potential privilege escalation vulnerabilities.

File System Permissions: Walks through the entire file system and checks if any files have incorrect permissions. Files with incorrect permissions can be exploited by attackers.

Remote Shell Connection: Attempts to establish an SSH connection to a specified target IP address. This tests for remote access vulnerabilities.

Firewall Test: Attempts to ping a specified target IP address. This tests for network connectivity and firewall rules.

To run:

Save the code as main.go in the perms_checker directory.
Build the binary: go build
Run the binary: ./perms_checker
Note:

Dummy File: Modify the dummyFile variable to match the actual path to your dummy file.
Remote IP: Replace the IP addresses in the net.ResolveTCPAddr and exec.Command calls with the actual target IP addresses.
Firewall Test: Adjust the exec.Command call to perform a different firewall test if desired (e.g., ICMP, TCP port scan).
User Permissions: This script requires root or equivalent privileges to perform all checks effectively.
Security: This script is for educational and testing purposes only. Use it responsibly and within the scope of authorized activities.
Disclaimer:
