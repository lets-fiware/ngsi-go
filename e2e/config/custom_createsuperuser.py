from django.contrib.auth.management.commands import createsuperuser

class Command(createsuperuser.Command):
    help = 'Create a superuser'

    def handle(self, *args, **options):
        options.setdefault('interactive', False)
        username = 'admin'
        email = 'admin@letsfiware.jp'
        password = '1234'
        database = options.get('database')

        user_data = {
            'username': username,
            'email': email,
            'password': password,
        }

        self.UserModel._default_manager.db_manager(database).create_superuser(**user_data)
