while true; do
    read -p "This will wipe ALL of your containers - make sure you know what you're doing. Continue? (y/n)" yn
    case $yn in
        [Yy]* )
            docker-compose down --volumes;
            docker system prune -a;

            docker-compose up --build;
            break;;
        [Nn]* )
            echo "Exiting script ..."
            exit;;
        * ) echo "Please answser y or n"
    esac 
done
    