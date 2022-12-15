#include "logger.h"

void info(string message) {
    cerr << "[INFO] " << message << endl;
}

void error(string message) {
    cerr << "[ERROR] " << message << " (" << __FILE__ << ":" << __LINE__ << ")\n";
    exit(1);
}
