my configuration for `squid` based caching server to prevent downloading RPM packages over and over again while playing around with virtual machines.

store_id_program
--------

using only the filename instead of the complete URL i can use alternate mirrors and still hit the cache, even if the RPM came originally from a different mirror.

Build:

    cd store_id_program && go build

there is a alternative store_id_program(store_id_program.py) provided by user [didib](https://github.com/didib), that takes the linux distribution into acount.

squid.conf
--------

i use the default configuration with the exception of refresh_patterns, which i replaced with my own:

    #                 3 month    12 month
    refresh_pattern . 129600 33% 525600

plus the additional configuration:

    cache_dir ufs /var/spool/squid 10000 16 256

    store_id_program /path/to/store_id_program
    store_id_children 5 startup=1

    # have not seen a larger RPM yet
    maximum_object_size 1 GB

    # cache RPMs only
    acl rpm_only urlpath_regex \.rpm
    cache allow rpm_only
    cache deny all
