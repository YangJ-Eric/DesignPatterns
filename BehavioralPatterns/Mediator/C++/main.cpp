#include <iostream>
#include <string>

/**
 * The Mediator interface declares a method used by components to notify the
 * mediator about various events. The Mediator may react to these events and
 * pass the execution to other components.
 */
class BaseComponent;
class Mediator {
public:
    virtual void Notify(BaseComponent *sender, std::string msg) const = 0;
};
/**
 * The Base Component provides the basic functionality of storing a mediator's
 * instance inside component objects.
 */
class BaseComponent{
protected:
    Mediator *mediator_;

public:
    BaseComponent(Mediator *mediator = nullptr) : mediator_(mediator){
    } 
    void set_mediator(Mediator *mediator){
        this->mediator_ = mediator;
    }
};

/**
 * Concrete Components implement various functionality. They don't depend on
 * other components. They also don't depend on any concrete mediator classes.
 */
class Component1 : public BaseComponent{
public:
    void SendMsg(std::string msg) {
        this->mediator_->Notify(this, msg);
    }
    void RecvMsg(std::string msg) {
        std::cout << "Component1 recv " << msg << "\n";
    }
};

class Component2 : public BaseComponent{
public:
    void SendMsg(std::string msg) {
        this->mediator_->Notify(this, msg);
    }
    void RecvMsg(std::string msg) {
        std::cout << "Component2 recv " << msg << "\n";
    }
};

/**
 * Concrete Mediators implement cooperative behavior by coordinating several
 * components.
 */
class ConcreteMediator : public Mediator
{
private:
    Component1 *component1_;
    Component2 *component2_;
public:
    ConcreteMediator(Component1 *c1, Component2 *c2) : component1_(c1), component2_(c2) {
        this->component1_->set_mediator(this);
        this->component2_->set_mediator(this);
    }
    void Notify(BaseComponent *sender, std::string msg) const override {
        if (sender == component1_){
            component2_->RecvMsg(msg);
        } else {
            component1_->RecvMsg(msg);
        } 
    }
};

int main()
{
    Component1 *c1 = new Component1;
    Component2 *c2 = new Component2;
    ConcreteMediator  *mediator = new ConcreteMediator(c1, c2);

    c1->SendMsg("I am Component1");
    c2->SendMsg("I am Component2");

    delete c1; 
    delete c2;
    delete mediator;
    return 0;
}