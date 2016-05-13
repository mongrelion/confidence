# confidence

### How to use this?
* Create your pool of servers on DigitalOcean with your desired configuration.

* Write your `hosts` file and assign the roles to each of the servers like this:

    ```
    uno  ansible_host=192.168.1.1
    dos  ansible_host=192.168.1.2
    tres ansible_host=192.168.1.3
    [appservers]
    uno
    dos
    [jenkins]
    tres
    ```

* Run the `bootstrap-local.yml` playbook so that you can get the servers' SSH fingerprint registered:

    ```
    $ ansible-playbook bootstrap-local.yml -i hosts
    ```

* Run the `bootstrap-remote.yml` playbook, which will install python in all of the servers so that you
  can later on run the rest of the playbooks in them, as Ansible requires all the servers to have it:

    ```
    $ ansible-playbook bootstrap-remote.yml -i hosts
    ```

* The `appservers.yml` playbook will download, compile and run the application on the app servers declared
  in the hosts file. Run it like this:

    ```
    $ ansible-playbook appservers.yml -i hosts
    ```

* The `jenkins.yml` playbook will install Jenkins on the servers with the jenkins role. Run it like this:

    ```
    $ ansible-playbook jenkins.yml -i hosts
    ```
