#! /bin/bash

# create by hetiulin
source /etc/profile
umask 0022
unset IFS
unset OFS
unset LD_PRELOAD
unset LD_LIBRARY_PATH
export PATH='/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin'

#if [ -w '/usr' ]; then
#    myPath="/usr/local/qcloud/stargate"
#else
#    myPath="/var/lib/qcloud/stargate"
#fi
#agent_name="$myPath/sgagent"

name=$1
tm=5
check_user()
{
    if [ "user_name" != "`whoami`" ]; then
        echo "Only user_name can execute this script"
        exit 2
    fi
}

check_alive()
{
    count=`ps -ef | grep "$name -d" | grep -v "grep" |wc -l`
    echo $count
    if [ $count -ge 1 ]; then
        # process exist
        echo "$name already exist"
    else
      $name -d
      echo "restart"
    fi
}

### Main Begin ###
check_user

ret=$?
if [ $ret -eq 0 ]; then
   for (( i = $tm; i<=60; i += $tm ))
     do
     #check_alive
      check_alive && sleep $tm && date
     done
fi
### Main End ###