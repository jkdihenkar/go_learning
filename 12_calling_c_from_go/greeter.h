#ifndef _GREETER_H
#define _GREETER_H

int greet(const char *name, int year, char *out);

struct Greetee {
    const char *name;
    int year;
};

int greet_struct(struct Greetee *g, char *out);

#endif