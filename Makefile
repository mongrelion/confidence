provision:
	@ansible-playbook -i hosts appservers.yml
	@ansible-playbook -i hosts jenkins.yml

clean:
	@rm *.retry
